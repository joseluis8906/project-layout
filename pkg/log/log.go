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
	reLevel   *regexp.Regexp
	reNoLevel *regexp.Regexp
)

func init() {
	reLevel = regexp.MustCompile(`^(?P<app>\w+) (?P<date>[0-9]{4}\/[0-9]{2}\/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}) (?P<line>\/[\w+-\/]+\.\w+:[0-9]+): (?P<level>INFO|ERROR) (?P<msg>.*)\n$`)
	reNoLevel = regexp.MustCompile(`^(?P<app>\w+) (?P<date>[0-9]{4}\/[0-9]{2}\/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}) (?P<line>\/[\w+-\/]+\.\w+:[0-9]+): (?P<msg>.*)\n$`)
}

func (l *Logger) Write(data []byte) (int, error) {
	if _, err := l.stdOut.Write(data); err != nil {
		return 0, err
	}

	var noLevel bool
	re := reLevel
	if !re.Match(data) {
		re = reNoLevel
		if !re.Match(data) {
			return 0, errors.New("log entry doesn't match")
		}

		noLevel = true
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
	line := slices.Index(groups, "line")
	var level int
	if !noLevel {
		level = slices.Index(groups, "level")
	}

	msg := slices.Index(groups, "msg")
	structured := make(map[string]string, 5)
	structured["app"] = values[app]
	structured["date"] = values[date]
	structured["line"] = values[line]
	if noLevel {
		structured["level"] = "ERROR"
	} else {
		structured["level"] = values[level]
	}

	structured["message"] = values[msg]
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
