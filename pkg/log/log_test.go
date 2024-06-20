package log

import (
	"os"
	"testing"
)

func TestLog_Write(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want error
	}{
		"with level": {
			in:   "delivery 2024/02/19 21:47:03 src/delivery/internal/storemanager/grpc/server.go:52: INFO calling storemanager.RegistersAStore: hello!",
			want: nil,
		},
		"without level": {
			in:   "delivery 2024/02/19 21:47:03 src/delivery/internal/storemanager/grpc/server.go:52: calling storemanager.RegistersAStore: hello!",
			want: nil,
		},
	}

	l := Logger{
		stdOut: os.Stderr,
	}

	for name, tc := range testCases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			length, err := l.Write([]byte(tc.in))
			if err != tc.want {
				t.Errorf("log.Write(%v) = %d, %v; want int, nil", tc.in, length, err)
			}
		})
	}

}
