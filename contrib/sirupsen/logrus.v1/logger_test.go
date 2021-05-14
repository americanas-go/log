package logrus

import (
	//"github.com/stretchr/testify/mock"
	"log/syslog"
	"reflect"
	"testing"

	"github.com/americanas-go/log"
	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"

	"github.com/stretchr/testify/suite"
)

type LoggerSuite struct {
	suite.Suite
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}

func (s *LoggerSuite) TestNewLogger() {

	tt := []struct {
		name string
		want func() log.Logger
		opts []Option
	}{
		{
			name: "New Logger with default options",
			want: func() log.Logger {
				return NewLoggerWithOptions(defaultOptions())
			},
			opts: []Option{},
		},
		{
			name: "New Logger with console enabled",
			want: func() log.Logger {
				opts := defaultOptions()
				opts.Console.Enabled = true
				return NewLoggerWithOptions(opts)
			},
			opts: []Option{
				WithConsoleEnabled(true),
			},
		},
		{
			name: "New Logger with console and file enabled",
			want: func() log.Logger {
				opts := defaultOptions()
				opts.Console.Enabled = true
				opts.File.Enabled = true
				return NewLoggerWithOptions(opts)
			},
			opts: []Option{
				WithConsoleEnabled(true),
				WithFileEnabled(true),
			},
		},
		{
			name: "New Logger with console disabled and file enabled",
			want: func() log.Logger {
				opts := defaultOptions()
				opts.Console.Enabled = false
				opts.File.Enabled = true
				return NewLoggerWithOptions(opts)
			},
			opts: []Option{
				WithConsoleEnabled(false),
				WithFileEnabled(true),
			},
		},
		{
			name: "New Logger with hook",
			want: func() log.Logger {
				opts := defaultOptions()
				opts.Hooks = []logrus.Hook{
					&logrus_syslog.SyslogHook{
						Writer: &syslog.Writer{},
					},
				}
				opts.Console.Enabled = false
				opts.File.Enabled = true
				return NewLoggerWithOptions(opts)
			},
			opts: []Option{
				WithHook(&logrus_syslog.SyslogHook{
					Writer: &syslog.Writer{},
				}),
				WithConsoleEnabled(false),
				WithFileEnabled(true),
			},
		},
	}

	for _, t := range tt {
		s.Run(t.name, func() {
			got := NewLogger(t.opts...)
			s.Assert().True(reflect.DeepEqual(got, t.want()), "NewLogger() = %v, want %v", got, t.want)
		})
	}
}

func (s *LoggerSuite) Test_logLevel() {
	tt := []struct {
		name  string
		level string
		want  logrus.Level
	}{
		{
			name:  "log level TRACE",
			level: "TRACE",
			want:  logrus.TraceLevel,
		},
		{
			name:  "log level DEBUG",
			level: "DEBUG",
			want:  logrus.DebugLevel,
		},
		{
			name:  "log level INFO",
			level: "INFO",
			want:  logrus.InfoLevel,
		},
		{
			name:  "log level ERROR",
			level: "ERROR",
			want:  logrus.ErrorLevel,
		},
		{
			name:  "log level WARN",
			level: "WARN",
			want:  logrus.WarnLevel,
		},
		{
			name:  "log level FATAL",
			level: "FATAL",
			want:  logrus.FatalLevel,
		},
		{
			name:  "log level PANIC",
			level: "PANIC",
			want:  logrus.PanicLevel,
		},
	}

	for _, t := range tt {
		s.Run(t.name, func() {
			got := logLevel(t.level)
			s.Assert().True(reflect.DeepEqual(got, t.want), "logLevel() = %v, want %v", got, t.want)
		})
	}
}
