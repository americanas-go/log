
log
=======

A simple, fast and consistent way for instantianting and using your favorite logging library in Golang. By a few changes in your config you can change the version or switch to a different library in seconds.

Installation
------------

	go get -u github.com/americanas-go/log


Supported libs
--------
* [Logrus](contrib/sirupsen/logrus.v1/README.md) - Is a structured logger for Go (golang), completely API compatible with the standard library logger.
* [Zap](contrib/go.uber.org/zap.v1/README.md) - Blazing fast, structured, leveled logging in Go.
* [Zerolog](contrib/rs/zerolog.v1/README.md) - Provides a fast and simple logger dedicated to JSON output.

Logger
--------
Logger is the contract for the logging.

#### func (*Logger) Printf
```go
func (l *Logger) Printf(format string, args ...interface{})
```
#### func (*Logger) Tracef
```go
func (l *Logger) Tracef(format string, args ...interface{})
```
#### func (*Logger) Trace
```go
func (l *Logger) Trace(args ...interface{})
```
#### func (*Logger) Debugf
```go
func (l *Logger) Debugf(format string, args ...interface{})
```
#### func (*Logger) Debug
```go
func (l *Logger) Debug(args ...interface{})
```
#### func (*Logger) Infof
```go
func (l *Logger) Infof(format string, args ...interface{})
```
#### func (*Logger) Warnf
```go
func (l *Logger) Warnf(format string, args ...interface{})
```
#### func (*Logger) Warn
```go
func (l *Logger) Warn(args ...interface{})
```
#### func (*Logger) Errorf
```go
func (l *Logger) Errorf(format string, args ...interface{})
```
#### func (*Logger) Error
```go
func (l *Logger) Error(args ...interface{})
```
#### func (*Logger) Fatalf
```go
func (l *Logger) Fatalf(format string, args ...interface{})
```
#### func (*Logger) Fatal
```go
func (l *Logger) Fatal(args ...interface{})
```
#### func (*Logger) Panicf
```go
func (l *Logger) Panicf(format string, args ...interface{})
```
#### func (*Logger) Panic
```go
func (l *Logger) Panic(args ...interface{})
```
#### func (*Logger) WithFields
```go
func (l *Logger) WithFields(f Fields) Logger
```
#### func (*Logger) WithField
```go
func (l *Logger) WithField(k string, v interface{}) Logger
```
#### func (*Logger) WithTypeOf
```go
func (l *Logger) WithTypeOf(obj interface{}) Logger
```
#### func (*Logger) ToContext
```go
func (l *Logger) ToContext(ctx context.Context) context.Context
```
#### func (*Logger) FromContext
```go
func (l *Logger) FromContext(ctx context.Context) Logger
```
#### func (*Logger) Output
```go
func (l *Logger) Output() io.Writer
```

Example
--------

```go
package main

import (
	"context"

	"github.com/americanas-go/log"

	//"github.com/americanas-go/log/contrib/go.uber.org/zap.v1"
	//"github.com/americanas-go/log/contrib/rs/zerolog.v1"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	ctx := context.Background()

	//logger := zap.NewLogger()
	//logger := zerolog.NewLogger()
	logger := logrus.NewLogger()

	logger = logger.WithField("main_field", "example")

	logger.Info("main method.")
	//zap output: 2021-05-16T14:30:31.788-0300	info	runtime/proc.go:225	main method.	{"main_field": "example"}
	//zerolog output: 2:30PM INF main method. main_field=example
	//logrus output: INFO[2021/05/16 14:31:12.477] main method.	main_field=example

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//zap output: 2021-05-16T14:30:31.788-0300	info	contrib/main.go:24	foo method.	{"main_field": "example", "foo_field": "example"}
	//zerolog output: 2:30PM INF foo method. foo_field=example main_field=example
	//logrus output: INFO[2021/05/16 14:31:12.477] foo method.	foo_field=example main_field=example

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//zap output: 2021-05-16T14:30:31.788-0300	info	contrib/main.go:37	bar method.	{"bar_field": "example", "main_field": "example", "foo_field": "example"}
	//zerolog output: 2:30PM INF bar method. bar_field=example foo_field=example main_field=example
	//logrus output: INFO[2021/05/16 14:31:12.477] bar method.	bar_field=example foo_field=example main_field=example
}

func withoutContext() {
	log.Info("withoutContext method")
	//zap output: 2021-05-16T14:30:31.788-0300	info	contrib/main.go:50	withoutContext method
	//zerolog output: 2:30PM INF withoutContext method
	//logrus output: INFO[2021/05/16 14:31:12.477] withoutContext method
}
```

Contributing
--------
Every help is always welcome. Fell free do throw us a pull request, we'll do our best to check it out as soon as possible. But before that, let us establish some guidelines:

1. This is an open source project so please do not add any proprietary code or infringe any copyright of any sort.
2. Avoid unnecessary dependencies or messing up go.mod file.
3. Be aware of golang coding style. Use a lint to help you out.
4.  Add tests to cover your contribution.
5. Use meaningful [messages](https://medium.com/@menuka/writing-meaningful-git-commit-messages-a62756b65c81) to your commits.
6. Use [pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests).
7. At last, but also important, be kind and polite with the community.

Any submitted issue which disrespect one or more guidelines above, will be discarded and closed.


<hr>

Released under the [MIT License](LICENSE).