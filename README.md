
log
=======

A simple, fast and consistent way for instantianting and using your favorite logging library in Golang. By a few changes in your config you can change the version or switch to a different library in seconds.

Installation
------------

	go get -u github.com/americanas-go/log


Supported libs
--------
* [Logrus](contrib/sirupsen/logrus.v1) - Is a structured logger for Go (golang), completely API compatible with the standard library logger.
* [Zap](contrib/go.uber.org/zap.v1) - Blazing fast, structured, leveled logging in Go.
* [Zerolog](contrib/rs/zerolog.v1) - Provides a fast and simple logger dedicated to JSON output.

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

Logger
--------
Logger is the contract for the logging.

#### func (*Logger) Printf
```go
func (l *Logger) Printf(format string, args ...interface{})
```
Printf logs a message at level Info (Logrus and Zap) and Debug (Zerolog) on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Printf("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Tracef
```go
func (l *Logger) Tracef(format string, args ...interface{})
```
Tracef logs a message at level Trace (Logrus and Zerolog) and Debug (Zap) on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Tracef("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Trace
```go
func (l *Logger) Trace(args ...interface{})
```
Trace logs a message at level Trace (Logrus and Zerolog) and Debug (Zap) on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Trace("hello world")
}
```
</details>
<br />

#### func (*Logger) Debugf
```go
func (l *Logger) Debugf(format string, args ...interface{})
```
Debugf logs a message at level Debug on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Debugf("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Debug
```go
func (l *Logger) Debug(args ...interface{})
```
Debug logs a message at level Debug on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Debug("hello world")
}
```
</details>
<br />

#### func (*Logger) Infof
```go
func (l *Logger) Infof(format string, args ...interface{})
```
Infof logs a message at level Info on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Infof("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Info
```go
func (l *Logger) Info(args ...interface{})
```
Info logs a message at level Info on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Info("hello world")
}
```
</details>
<br />

#### func (*Logger) Warnf
```go
func (l *Logger) Warnf(format string, args ...interface{})
```
Warnf logs a message at level Warn on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Warnf("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Warn
```go
func (l *Logger) Warn(args ...interface{})
```
Warn logs a message at level Warn on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Warn("hello world")
}
```
</details>
<br />

#### func (*Logger) Errorf
```go
func (l *Logger) Errorf(format string, args ...interface{})
```
Errorf logs a message at level Error on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Errorf("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Error
```go
func (l *Logger) Error(args ...interface{})
```
Error logs a message at level Error on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Error("hello world")
}
```
</details>
<br />

#### func (*Logger) Fatalf
```go
func (l *Logger) Fatalf(format string, args ...interface{})
```
Fatalf logs a message at level Fatal on the standard logger, then calls os.Exit(1).
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Fatalf("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Fatal
```go
func (l *Logger) Fatal(args ...interface{})
```
Fatal logs a message at level Fatal on the standard logger, then calls os.Exit(1).
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Fatal("hello world")
}
```
</details>
<br />

#### func (*Logger) Panicf
```go
func (l *Logger) Panicf(format string, args ...interface{})
```
Panicf logs a message at level Panic on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Panicf("hello %s", "world")
}
```
</details>
<br />

#### func (*Logger) Panic
```go
func (l *Logger) Panic(args ...interface{})
```
Panic logs a message at level Panic on the standard logger.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger.Panic("hello world")
}
```
</details>
<br />

#### func (*Logger) WithFields
```go
func (l *Logger) WithFields(f Fields) Logger
```
WithFields creates an entry from the standard logger and adds multiple fields to i
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger = logger.WithFields(log.Fields{
		"hello": "world",
		"foo":   "bar",
	})
	logger.Info("main method.")
}
```
</details>
<br />

#### func (*Logger) WithField
```go
func (l *Logger) WithField(k string, v interface{}) Logger
```
WithField creates an entry from the standard logger and adds a field to it.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	logger = logger.WithField("hello", "world")
	logger.Info("main method.")
}
```
</details>
<br />

#### func (*Logger) WithTypeOf
```go
func (l *Logger) WithTypeOf(obj interface{}) Logger
```
WithTypeOf creates an entry from the standard logger and adds type and package information fields.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	s := S{
		F0: "foo",
		F1: "bar",
	}
	logger = logger.WithTypeOf(s)
	logger.Info("main method.")
}

type S struct {
	F0 string `alias:"field_0"`
	F1 string `alias:""`
}
```
</details>
<br />

#### func (*Logger) ToContext
```go
func (l *Logger) ToContext(ctx context.Context) context.Context
```
ToContext returns a copy of ctx in which its fields are added to those of l.
<details>
<summary>Example</summary>

```go
package main

import (
	"context"

	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	ctx := context.Background()

	logger := logrus.NewLogger()
	logger = logger.WithField("main_field", "example")
	logger.Info("main method.")

	ctx = logger.ToContext(ctx)
	logger.Infof("%s method.", "main")
}
```
</details>
<br />

#### func (*Logger) FromContext
```go
func (l *Logger) FromContext(ctx context.Context) Logger
```
FromContext returns a Logger from ctx.
<details>
<summary>Example</summary>

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	ctx := context.Background()

	logger := logrus.NewLogger()
	logger = logger.WithField("main_field", "example")
	logger.Info("main method.")

	foo(ctx)
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Infof("%s method.", "foo")
}
```
</details>
<br />

#### func (*Logger) Output
```go
func (l *Logger) Output() io.Writer
```
Output returns a io.Writer.
<details>
<summary>Example</summary>

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logger := logrus.NewLogger()
	n, err := logger.Output().Write([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	logger.Infof("%d bytes", n)
}
```
</details>
<br />

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