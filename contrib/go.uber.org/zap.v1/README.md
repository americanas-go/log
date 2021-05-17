zap.v1
=======

Examples
--------
### Simple Logging Example

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/go.uber.org/zap.v1"
)

func main() {
	ctx := context.Background()

	//example use zap
	logger := zap.NewLogger()

	logger = logger.WithField("main_field", "example")

	logger.Info("main method.")
	//output: 2021-05-17T15:22:04.126-0300	info	runtime/proc.go:225	main method.	{"main_field": "example"}

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//output: 2021-05-17T15:22:04.126-0300	info	contrib/main.go:24	foo method.	{"main_field": "example", "foo_field": "example"}

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//output: 2021-05-17T15:22:04.126-0300	info	contrib/main.go:37	bar method.	{"bar_field": "example", "main_field": "example", "foo_field": "example"}
}

func withoutContext() {
	log.Info("withoutContext method")
	//output: 2021-05-17T15:22:04.126-0300	info	contrib/main.go:50	withoutContext method
}
```

### Logging With Options Example

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/go.uber.org/zap.v1"
)

func main() {
	ctx := context.Background()

	//example use zap with options
	logger := zap.NewLogger(withConsoleFormatter(), withConsoleLevel())

	logger = logger.WithField("main_field", "example")

	logger.Info("main method.")
	//output: {"level":"info","ts":"2021-05-17T15:26:57.624-0300","caller":"runtime/proc.go:225","msg":"main method.","main_field":"example"}

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//output: {"level":"info","ts":"2021-05-17T15:26:57.624-0300","caller":"contrib/main.go:23","msg":"foo method.","main_field":"example","foo_field":"example"}

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//output: {"level":"info","ts":"2021-05-17T15:26:57.624-0300","caller":"contrib/main.go:36","msg":"bar method.","main_field":"example","foo_field":"example","bar_field":"example"}
}

func withoutContext() {
	log.Info("withoutContext method")
	//output: {"level":"info","ts":"2021-05-17T15:26:57.624-0300","caller":"contrib/main.go:49","msg":"withoutContext method"}
}

func withConsoleFormatter() zap.Option {
	return func(o *zap.Options) {
		o.Console.Formatter = "JSON"
	}
}

func withConsoleLevel() zap.Option {
	return func(o *zap.Options) {
		o.Console.Level = "DEBUG"
	}
}
```
