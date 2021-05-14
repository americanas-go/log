package log_test

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func ExampleNewLogger() {
	log.NewLogger(logrus.NewLogger())
	log.WithField("main_field", "example")
	log.Info("main method.")
	// Output: INFO[2021/05/14 17:15:04.757] main method. main_field=example
}

func ExampleToContext() {
	bar := func(ctx context.Context) {
		logger := log.FromContext(ctx)

		logger = logger.WithField("bar_field", "example")
		logger.Infof("%s method.", "bar")
	}

	foo := func(ctx context.Context) {
		logger := log.FromContext(ctx)

		logger.WithField("foo_field", "example")
		logger.Infof("%s method.", "foo")

		ctx = logger.ToContext(ctx)

		bar(ctx)
	}

	withoutContext := func() {
		log.Info("withoutContext method")
	}

	ctx := context.Background()
	log.NewLogger(logrus.NewLogger())

	ctx = log.ToContext(ctx)

	foo(ctx)

	withoutContext()

	// Output:
	// INFO[2021/05/14 17:15:04.757] foo method. foo_field=example main_field=example
	// INFO[2021/05/14 17:15:04.757] bar method. bar_field=example foo_field=example main_field=example
	// INFO[2021/05/14 17:15:04.757] withoutContext method
}

func ExampleFromContext() {
	bar := func(ctx context.Context) {
		logger := log.FromContext(ctx)

		logger = logger.WithField("bar_field", "example")
		logger.Infof("%s method.", "bar")
	}

	foo := func(ctx context.Context) {
		logger := log.FromContext(ctx)

		logger.WithField("foo_field", "example")
		logger.Infof("%s method.", "foo")

		ctx = logger.ToContext(ctx)

		bar(ctx)
	}

	withoutContext := func() {
		log.Info("withoutContext method")
	}

	ctx := context.Background()
	log.NewLogger(logrus.NewLogger())

	ctx = log.ToContext(ctx)

	foo(ctx)

	withoutContext()

	// Output:
	// INFO[2021/05/14 17:15:04.757] foo method. foo_field=example main_field=example
	// INFO[2021/05/14 17:15:04.757] bar method. bar_field=example foo_field=example main_field=example
	// INFO[2021/05/14 17:15:04.757] withoutContext method
}
