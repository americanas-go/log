zerolog.v1
=======

Examples
--------
### Simple Logging Example

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

func main() {
	ctx := context.Background()

	//example use zerolog
	logger := zerolog.NewLogger()

	logger = logger.WithField("main_field", "example")

	logger.Info("main method.")
	//output: 3:34PM INF main method. main_field=example

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//output: 3:34PM INF foo method. foo_field=example main_field=example

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//output: 3:34PM INF bar method. bar_field=example foo_field=example main_field=example
}

func withoutContext() {
	log.Info("withoutContext method")
	//output: 3:34PM INF withoutContext method
}
```

### Logging With Options Example

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

func main() {
	ctx := context.Background()

	//example use zerolog with options
	logger := zerolog.NewLogger(withFormatter(), withConsoleLevel())

	logger = logger.WithField("main_field", "example")

	logger.Info("main method.")
	//output: {"log_level":"info","main_field":"example","time":"2021-05-17T15:32:20-03:00","log_message":"main method."}

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//output: {"log_level":"info","main_field":"example","foo_field":"example","time":"2021-05-17T15:32:20-03:00","log_message":"foo method."}

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//output: {"log_level":"info","main_field":"example","foo_field":"example","bar_field":"example","time":"2021-05-17T15:32:20-03:00","log_message":"bar method."}
}

func withoutContext() {
	log.Info("withoutContext method")
	//output: {"log_level":"info","time":"2021-05-17T15:32:20-03:00","log_message":"withoutContext method"}
}

func withFormatter() zerolog.Option {
	return func(o *zerolog.Options) {
		o.Formatter = "JSON"
	}
}

func withConsoleLevel() zerolog.Option {
	return func(o *zerolog.Options) {
		o.Console.Level = "DEBUG"
	}
}
```
