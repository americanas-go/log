package zerolog

import (
	"bytes"
	"context"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/americanas-go/log"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ctxKey string

const (
	key                   ctxKey = "ctxfields"
	defaultFormatter             = "TEXT"
	defaultConsoleEnabled        = true
	defaultConsoleLevel          = "INFO"
	defaultFileEnabled           = false
	defaultFileLevel             = "INFO"
	defaultFilePath              = "/tmp"
	defaultFileName              = "application.log"
	defaultFileMaxSize           = 100
	defaultFileCompress          = true
	defaultFileMaxAge            = 28
)

func NewLogger(option ...Option) log.Logger {
	options := options(option)
	return NewLoggerWithOptions(options)
}

func NewLoggerWithOptions(options *Options) log.Logger {
	writer := getWriter(options)
	if writer == nil {
		zerologger := zerolog.Nop()
		logger := &logger{
			logger: zerologger,
		}

		log.NewLogger(logger)
		return logger
	}

	zerolog.MessageFieldName = "log_message"
	zerolog.LevelFieldName = "log_level"

	zerologger := zerolog.New(writer).With().Timestamp().Logger()
	level := getLogLevel(options.Console.Level)
	zerologger = zerologger.Level(level)

	logger := &logger{
		logger: zerologger,
		writer: writer,
		fields: log.Fields{},
	}

	log.NewLogger(logger)
	return logger
}

func defaultOptions() *Options {
	return &Options{
		Formatter: defaultFormatter,
		Console: struct {
			Enabled bool
			Level   string
		}{
			Enabled: defaultConsoleEnabled,
			Level:   defaultConsoleLevel,
		},
		File: struct {
			Enabled  bool
			Level    string
			Path     string
			Name     string
			MaxSize  int
			Compress bool
			MaxAge   int
		}{
			Enabled:  defaultFileEnabled,
			Level:    defaultFileLevel,
			Path:     defaultFilePath,
			Name:     defaultFileName,
			MaxSize:  defaultFileMaxSize,
			Compress: defaultFileCompress,
			MaxAge:   defaultFileMaxAge,
		},
	}
}

func options(option []Option) *Options {
	options := defaultOptions()

	for _, o := range option {
		o(options)
	}
	return options
}

type logger struct {
	logger zerolog.Logger
	writer io.Writer
	fields log.Fields
}

func getLogLevel(level string) zerolog.Level {
	switch level {
	case "DEBUG":
		return zerolog.DebugLevel
	case "WARN":
		return zerolog.WarnLevel
	case "FATAL":
		return zerolog.FatalLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "TRACE":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}

func getWriter(options *Options) io.Writer {
	var writer io.Writer
	switch options.Formatter {
	case "TEXT":
		writer = zerolog.ConsoleWriter{Out: os.Stdout}
	default:
		writer = os.Stdout
	}

	if options.File.Enabled {
		s := []string{options.File.Path, "/", options.File.Name}
		fileLocation := strings.Join(s, "")

		fileHandler := &lumberjack.Logger{
			Filename: fileLocation,
			MaxSize:  options.File.MaxSize,
			Compress: options.File.Compress,
			MaxAge:   options.File.MaxAge,
		}

		if options.Console.Enabled {
			return io.MultiWriter(writer, fileHandler)
		}
		return fileHandler
	} else if options.Console.Enabled {
		return writer
	}

	return nil
}

func (l *logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.logger.Trace().Msgf(format, args...)
}

func (l *logger) Trace(args ...interface{}) {
	format := bytes.NewBufferString("")
	for range args {
		format.WriteString("%v")
	}

	l.logger.Trace().Msgf(format.String(), args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	format := bytes.NewBufferString("")
	for range args {
		format.WriteString("%v")
	}

	l.logger.Debug().Msgf(format.String(), args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	format := bytes.NewBufferString("")
	for range args {
		format.WriteString("%v")
	}

	l.logger.Info().Msgf(format.String(), args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	format := bytes.NewBufferString("")
	for range args {
		format.WriteString("%v")
	}

	l.logger.Warn().Msgf(format.String(), args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	format := bytes.NewBufferString("")
	for range args {
		format.WriteString("%v")
	}

	l.logger.Error().Msgf(format.String(), args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	format := bytes.NewBufferString("")
	for range args {
		format.WriteString("%v")
	}

	l.logger.Fatal().Msgf(format.String(), args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

func (l *logger) Panic(args ...interface{}) {
	format := bytes.NewBufferString("")
	for range args {
		format.WriteString("%v")
	}

	l.logger.Panic().Msgf(format.String(), args...)
}

func (l *logger) WithField(key string, value interface{}) log.Logger {
	newField := log.Fields{}
	newField[key] = value

	newLogger := l.logger.With().Fields(newField).Logger()
	return &logger{newLogger, l.writer, newField}
}

func (l *logger) WithFields(fields log.Fields) log.Logger {
	newLogger := l.logger.With().Fields(fields).Logger()
	return &logger{newLogger, l.writer, fields}
}

func (l *logger) WithTypeOf(obj interface{}) log.Logger {

	t := reflect.TypeOf(obj)

	return l.WithFields(log.Fields{
		"reflect.type.name":    t.Name(),
		"reflect.type.package": t.PkgPath(),
	})
}

func (l *logger) Fields() log.Fields {
	return l.fields
}

func (l *logger) Output() io.Writer {
	return l.writer
}

func (l *logger) ToContext(ctx context.Context) context.Context {
	logger := l.logger
	return logger.WithContext(ctx)
}

func (l *logger) FromContext(ctx context.Context) log.Logger {
	zerologger := zerolog.Ctx(ctx)
	if zerologger.GetLevel() == zerolog.Disabled {
		return l
	}
	rawFields := ctx.Value(key)
	fields := log.Fields{}
	if rawFields != nil {
		switch v := rawFields.(type) {
		case log.Fields:
			fields = v
		}
	}
	return &logger{*zerologger, l.writer, fields}
}
