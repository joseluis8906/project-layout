package log

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type (
	Deps struct {
		fx.In
		Config *viper.Viper
	}

	Logger struct {
		stdOut io.Writer
		conn   *fluent.Fluent
		tag    string
	}
)

var (
	re        *regexp.Regexp
	reNoLevel *regexp.Regexp
)

func init() {
	re = regexp.MustCompile(`^(?P<app>\w+) (?P<date>[0-9]{4}\/[0-9]{2}\/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}) (?P<line>\/*[\w+-\/]+\.\w+:[0-9]+): (?P<level>[INFO|ERROR]*)(?P<msg>.*)\n$`)
}

func (l *Logger) Write(data []byte) (int, error) {
	if _, err := l.stdOut.Write(data); err != nil {
		return 0, err
	}

	re := re
	if !re.Match(data) {
		return 0, errors.New("log entry doesn't match")
	}

	var transformer strings.Builder
	transformer.Grow(len(data))
	for _, b := range data {
		fmt.Fprintf(&transformer, "%c", b)
	}

	wholeMsj := transformer.String()
	groups := re.SubexpNames()
	values := re.FindStringSubmatch(wholeMsj)
	app := slices.Index(groups, "app")
	date := slices.Index(groups, "date")
	level := slices.Index(groups, "level")
	line := slices.Index(groups, "line")

	msg := slices.Index(groups, "msg")
	structured := make(map[string]string, 5)
	structured["app"] = values[app]
	structured["date"] = values[date]
	structured["line"] = values[line]
	structured["level"] = values[level]
	if len(structured["level"]) == 0 {
		structured["level"] = "ERROR"
	}
	structured["message"] = values[msg]
	if l.conn == nil {
		return 0, errors.New("fluentd connection is nil")
	}

	return len(data), l.conn.Post(l.tag, structured)
}

func New(deps Deps) *log.Logger {
	fluentd, err := fluent.New(fluent.Config{
		FluentHost:    deps.Config.GetString("fluentd.host"),
		FluentPort:    deps.Config.GetInt("fluentd.port"),
		FluentNetwork: "tcp",
		MarshalAsJSON: true,
		Async:         true,
	})

	if err != nil {
		log.Fatalf("connecting fluentd: %v", err)
	}

	logger := Logger{
		stdOut: os.Stderr,
		conn:   fluentd,
		tag:    deps.Config.GetString("fluentd.tag"),
	}

	l := log.New(&logger, fmt.Sprintf("%s ", deps.Config.GetString("app.name")), log.LstdFlags)
	l.SetFlags(log.LstdFlags | log.Llongfile)
	return l
}

func Info(message string) string {
	return fmt.Sprintf("INFO %v", message)
}
