package text

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

type Option func(formatter *logrus.TextFormatter)

func New(options ...Option) logrus.Formatter {

	fmt := &logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006/01/02 15:04:05.000",
		DisableSorting:            false,
		DisableLevelTruncation:    true,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
	}

	for _, option := range options {
		option(fmt)
	}

	return fmt
}

func WithForceColors(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.ForceColors = value
	}
}

func WithDisableColors(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableColors = value
	}
}

func WithForceQuote(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.ForceQuote = value
	}
}

func WithDisableQuote(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableQuote = value
	}
}

func WithEnvironmentOverrideColors(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.EnvironmentOverrideColors = value
	}
}

func WithDisableTimestamp(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableTimestamp = value
	}
}

func WithFullTimestamp(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.FullTimestamp = value
	}
}

func WithTimestampFormat(value string) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.TimestampFormat = value
	}
}

func WithDisableSorting(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableSorting = value
	}
}

func WithDisableLevelTruncation(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.DisableLevelTruncation = value
	}
}

func WithPadLevelText(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.PadLevelText = value
	}
}

func WithQuoteEmptyFields(value bool) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.QuoteEmptyFields = value
	}
}

func WithFieldMap(value logrus.FieldMap) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.FieldMap = value
	}
}

func WithCallerPrettyfier(value func(*runtime.Frame) (function string, file string)) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.CallerPrettyfier = value
	}
}

func WithSortingFunc(value func([]string)) Option {
	return func(formatter *logrus.TextFormatter) {
		formatter.SortingFunc = value
	}
}
