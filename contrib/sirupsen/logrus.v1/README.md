logrus.v1
=======

Examples
--------
### Simple Logging Example

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	ctx := context.Background()

	//example use logrus
	logger := logrus.NewLogger()

	logger = logger.WithField("main_field", "example")

	logger.Info("main method.")
	//output: INFO[2021/05/14 17:15:04.757] main method. main_field=example

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//output: INFO[2021/05/14 17:15:04.757] foo method. foo_field=example main_field=example

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//output: INFO[2021/05/14 17:15:04.757] bar method. bar_field=example foo_field=example main_field=example
}

func withoutContext() {
	log.Info("withoutContext method")
	//output: INFO[2021/05/14 17:15:04.757] withoutContext method
}
```

### Logging With Options Example

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1/formatter/json"
)

func main() {
	ctx := context.Background()

	//example use logrus with options
	logger := logrus.NewLogger(withFormatter(), withConsoleLevel())

	logger = logger.WithField("main_field", "example")

	logger.Info("main method.")
	//output: {"level":"info","main_field":"example","msg":"main method.","time":"2021/05/17 14:16:12.514"}

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//output: {"foo_field":"example","level":"info","main_field":"example","msg":"foo method.","time":"2021/05/17 14:16:12.514"}

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//output: {"bar_field":"example","foo_field":"example","level":"info","main_field":"example","msg":"bar method.","time":"2021/05/17 14:16:12.514"}
}

func withoutContext() {
	log.Info("withoutContext method")
	//output: {"level":"info","msg":"withoutContext method","time":"2021/05/17 14:16:12.514"}
}

func withFormatter() logrus.Option {
	return func(o *logrus.Options) {
		o.Formatter = json.New()
	}
}

func withConsoleLevel() logrus.Option {
	return func(o *logrus.Options) {
		o.Console.Level = "DEBUG"
	}
}
```
