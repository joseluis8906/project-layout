package log

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestLog_Write(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want error
	}{
		"with level": {
			in:   "app 2024/02/19 21:47:03 internal/app/grpc/server.go:52: INFO calling app.DoSomething: the msg!\n",
			want: errors.New("fluentd connection is nil"),
		},
		"without level": {
			in:   "app 2024/02/19 21:47:03 internal/app/grpc/server.go:52: calling app.DoSomething: the msg!\n",
			want: errors.New("fluentd connection is nil"),
		},
	}

	l := Logger{
		stdOut: os.Stderr,
	}
	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			length, err := l.Write([]byte(tc.in))
			if fmt.Sprintf("%s", err) != fmt.Sprintf("%s", tc.want) {
				t.Errorf("log.Write(%q) = %d, %v; want int, <nil>", tc.in, length, err)
			}
		})
	}

}
