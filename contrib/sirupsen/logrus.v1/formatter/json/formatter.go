package json

import "github.com/sirupsen/logrus"

type Option func(formatter *logrus.JSONFormatter)

func New(options ...Option) logrus.Formatter {
	fmt := &logrus.JSONFormatter{
		TimestampFormat:   "2006/01/02 15:04:05.000",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		PrettyPrint:       false,
	}

	for _, option := range options {
		option(fmt)
	}

	return fmt
}

func WithTimestampFormat(value string) Option {
	return func(formatter *logrus.JSONFormatter) {
		formatter.TimestampFormat = value
	}
}

func WithDisableTimestamp(value bool) Option {
	return func(formatter *logrus.JSONFormatter) {
		formatter.DisableTimestamp = value
	}
}

func WithDisableHTMLEscape(value bool) Option {
	return func(formatter *logrus.JSONFormatter) {
		formatter.DisableHTMLEscape = value
	}
}

func WithDataKey(value string) Option {
	return func(formatter *logrus.JSONFormatter) {
		formatter.DataKey = value
	}
}

func WithFieldMap(value logrus.FieldMap) Option {
	return func(formatter *logrus.JSONFormatter) {
		formatter.FieldMap = value
	}
}

func WithPrettyPrint(value bool) Option {
	return func(formatter *logrus.JSONFormatter) {
		formatter.PrettyPrint = value
	}
}
