package logrus

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/americanas-go/log"
	logredis "github.com/jpfaria/logrus-redis-hook"
	"github.com/ravernkoh/cwlogsfmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ctxKey string

const key ctxKey = "ctxfields"

func NewLoggerWithFormatter(formatter logrus.Formatter, options *Options) log.Logger {

	lLogger := new(logrus.Logger)

	if options.Redis.Enabled {

		hookConfig := logredis.HookConfig{
			Host:   options.Redis.Host,
			Key:    options.Redis.Key,
			Format: options.Redis.Format,
			App:    options.Redis.App,
			Port:   options.Redis.Port,
			DB:     options.Redis.DB,
		}

		hook, err := logredis.NewHook(hookConfig)
		if err == nil {
			lLogger.AddHook(hook)
		} else {
			lLogger.Errorf("logredis error: %q", err)
		}

	}

	var fileHandler *lumberjack.Logger

	lLogger.SetOutput(ioutil.Discard)

	if options.File.Enabled {

		s := []string{options.File.Path, "/", options.File.Name}
		fileLocation := strings.Join(s, "")

		fileHandler = &lumberjack.Logger{
			Filename: fileLocation,
			MaxSize:  options.File.MaxSize,
			Compress: options.File.Compress,
			MaxAge:   options.File.MaxAge,
		}

	}

	if options.Console.Enabled && options.File.Enabled {
		lLogger.SetOutput(io.MultiWriter(os.Stdout, fileHandler))
	} else if options.File.Enabled {
		lLogger.SetOutput(fileHandler)
	} else if options.Console.Enabled {
		lLogger.SetOutput(os.Stdout)
	}

	level := getLogLevel(options.Console.Level)
	lLogger.SetLevel(level)

	lLogger.SetFormatter(formatter)

	logger := &logger{
		logger: lLogger,
	}

	log.NewLogger(logger)
	return logger

}

func NewLogger(options *Options) log.Logger {
	formatter := getFormatter(options)
	return NewLoggerWithFormatter(formatter, options)
}

func getLogLevel(level string) logrus.Level {

	switch level {

	case "DEBUG":
		return logrus.DebugLevel
	case "WARN":
		return logrus.WarnLevel
	case "FATAL":
		return logrus.FatalLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "TRACE":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}

}

func getFormatter(options *Options) logrus.Formatter {

	var formatter logrus.Formatter

	switch options.Formatter {

	case "JSON":

		fmt := &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "date",
				logrus.FieldKeyLevel: "log_level",
				logrus.FieldKeyMsg:   "log_message",
			},
		}

		fmt.TimestampFormat = options.Time.Format

		formatter = fmt

	case "AWS_CLOUD_WATCH":

		formatter = &cwlogsfmt.CloudWatchLogsFormatter{
			PrefixFields:     []string{"RequestId"},
			QuoteEmptyFields: true,
		}

	default:

		fmt := &logrus.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
		}
		fmt.TimestampFormat = options.Time.Format

		formatter = fmt

	}

	return formatter
}

type logger struct {
	logger *logrus.Logger
	fields log.Fields
}

func (l *logger) Printf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}

func (l *logger) Tracef(format string, args ...interface{}) {
	l.logger.Tracef(format, args...)
}

func (l *logger) Trace(args ...interface{}) {
	l.logger.Trace(args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *logger) WithField(key string, value interface{}) log.Logger {

	entry := l.logger.WithField(key, value)

	return &logEntry{
		entry:  entry,
		fields: convertToFields(entry.Data),
	}
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logger) WithFields(fields log.Fields) log.Logger {
	return &logEntry{
		entry:  l.logger.WithFields(convertToLogrusFields(fields)),
		fields: fields,
	}
}

func (l *logger) WithTypeOf(obj interface{}) log.Logger {

	t := reflect.TypeOf(obj)

	return l.WithFields(log.Fields{
		"reflect.type.name":    t.Name(),
		"reflect.type.package": t.PkgPath(),
	})
}

func (l *logger) GetFields() log.Fields {
	return l.fields
}

func (l *logger) Output() io.Writer {
	return l.logger.Out
}

func (l *logger) ToContext(ctx context.Context) context.Context {
	return toContext(ctx, l.fields)
}

func (l *logger) FromContext(ctx context.Context) log.Logger {
	fields := fieldsFromContext(ctx)
	return l.WithFields(fields)
}

type logEntry struct {
	entry  *logrus.Entry
	fields log.Fields
}

func (l *logEntry) Printf(format string, args ...interface{}) {
	l.entry.Printf(format, args...)
}

func (l *logEntry) Tracef(format string, args ...interface{}) {
	l.entry.Tracef(format, args...)
}

func (l *logEntry) Trace(args ...interface{}) {
	l.entry.Trace(args...)
}

func (l *logEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *logEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *logEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *logEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *logEntry) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *logEntry) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

func (l *logEntry) WithField(key string, value interface{}) log.Logger {

	entry := l.entry.WithField(key, value)

	return &logEntry{
		entry:  entry,
		fields: convertToFields(entry.Data),
	}
}

func (l *logEntry) Output() io.Writer {
	return l.entry.Logger.Out
}

func (l *logEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logEntry) Panicf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logEntry) WithFields(fields log.Fields) log.Logger {
	return &logEntry{
		entry: l.entry.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logEntry) WithTypeOf(obj interface{}) log.Logger {

	t := reflect.TypeOf(obj)

	return l.WithFields(log.Fields{
		"reflect.type.name":    t.Name(),
		"reflect.type.package": t.PkgPath(),
	})
}

func (l *logEntry) ToContext(ctx context.Context) context.Context {
	return toContext(ctx, l.fields)
}

func (l *logEntry) FromContext(ctx context.Context) log.Logger {
	fields := fieldsFromContext(ctx)
	return l.WithFields(fields)
}

func toContext(ctx context.Context, fields log.Fields) context.Context {
	ctxFields := fieldsFromContext(ctx)

	if ctxFields == nil {
		ctxFields = map[string]interface{}{}
	}

	for k, v := range fields {
		ctxFields[k] = v
	}

	return context.WithValue(ctx, key, ctxFields)
}

func fieldsFromContext(ctx context.Context) log.Fields {
	fields := make(log.Fields)

	if ctx == nil {
		return fields
	}

	if f, ok := ctx.Value(key).(log.Fields); ok && f != nil {
		for k, v := range f {
			fields[k] = v
		}
	}

	return fields
}

func convertToLogrusFields(fields log.Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}

func convertToFields(logrusFields logrus.Fields) log.Fields {
	fields := log.Fields{}
	for index, val := range logrusFields {
		fields[index] = val
	}
	return fields
}
