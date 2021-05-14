package cloudwatch

import (
	"github.com/ravernkoh/cwlogsfmt"
	"github.com/sirupsen/logrus"
)

type Option func(formatter *cwlogsfmt.CloudWatchLogsFormatter)

func NewFormatter(options ...Option) logrus.Formatter {
	fmt := &cwlogsfmt.CloudWatchLogsFormatter{
		PrefixFields:     []string{"RequestId"},
		DisableSorting:   false,
		QuoteEmptyFields: true,
	}

	for _, option := range options {
		option(fmt)
	}

	return fmt
}

func WithPrefixFields(value []string) Option {
	return func(formatter *cwlogsfmt.CloudWatchLogsFormatter) {
		formatter.PrefixFields = value
	}
}

func WithDisableSorting(value bool) Option {
	return func(formatter *cwlogsfmt.CloudWatchLogsFormatter) {
		formatter.DisableSorting = value
	}
}

func WithQuoteEmptyFields(value bool) Option {
	return func(formatter *cwlogsfmt.CloudWatchLogsFormatter) {
		formatter.QuoteEmptyFields = value
	}
}
