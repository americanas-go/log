
log
=======
[![Go Reference](https://pkg.go.dev/badge/github.com/americanas-go/log.svg)](https://pkg.go.dev/github.com/americanas-go/log)
[![changelog](https://camo.githubusercontent.com/4d89fc2186d69bdbb2c6ea6cb54ab16915be5e5e0b63a393e87a75741f1baa8c/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6368616e67656c6f672d4348414e47454c4f472e6d642d253233453035373335)](CHANGELOG.md)

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

Global Logger
--------
The `americanas-go/log` provides top level logging function, however by default they do nothing (NoOp). You can define your global logger, after you instantiate the desired implementation, by using the `log.SetGlobalLogger`.

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
	//logger := zap.NewLogger()
	//logger := zerolog.NewLogger()
	logger := logrus.NewLogger()
	log.SetGlobalLogger(logger)

	log.Info("main method.")
}
```


Logger
--------
Logger is the contract for the logging.

#### Printf
logs a message at level Info (Logrus and Zap) and Debug (Zerolog) on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Printf("hello %s", "world")
}
```

#### Tracef
logs a message at level Trace (Logrus and Zerolog) and Debug (Zap) on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Tracef("hello %s", "world")
}
```

#### Trace
logs a message at level Trace (Logrus and Zerolog) and Debug (Zap) on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Trace("hello world")
}
```

#### Debugf
logs a message at level Debug on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Debugf("hello %s", "world")
}
```

#### Debug
logs a message at level Debug on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Debug("hello world")
}
```

#### Infof
logs a message at level Info on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Infof("hello %s", "world")
}
```

#### Info
logs a message at level Info on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Info("hello world")
}
```

#### Warnf
logs a message at level Warn on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Warnf("hello %s", "world")
}
```

#### Warn
logs a message at level Warn on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Warn("hello world")
}
```

#### Errorf
logs a message at level Error on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Errorf("hello %s", "world")
}
```

#### Error
logs a message at level Error on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Error("hello world")
}
```

#### Fatalf
logs a message at level Fatal on the standard logger, then calls os.Exit(1).

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Fatalf("hello %s", "world")
}
```

#### Fatal
logs a message at level Fatal on the standard logger, then calls os.Exit(1).

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Fatal("hello world")
}
```

#### Panicf
logs a message at level Panic on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Panicf("hello %s", "world")
}
```

#### Panic
logs a message at level Panic on the standard logger.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.Panic("hello world")
}
```

#### WithFields
creates an entry from the standard logger and adds multiple fields to it.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {

	log.SetGlobalLogger(logrus.NewLogger())

	log.WithFields(log.Fields{
		"hello": "world",
		"foo":   "bar",
	}).Info("main method.")
}
```

#### WithField
creates an entry from the standard logger and adds a field to it.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	
	log.WithField("hello", "world").
		Info("main method.")
}
```

#### WithTypeOf
creates an entry from the standard logger and adds type and package information to it.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	
	s := S{}
	s.Foo()
}

type S struct {}

func (s *S) Foo() {
	logger := log.WithTypeOf(s)
	logger.Info("main method.")
}
```

#### WithError
creates an entry from the standard logger with the error content as a field.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	log.SetGlobalLogger(logrus.NewLogger())
	log.WithError(errors.New("something bad")).
		Info("main method.")
}
```

#### ToContext/FromContext
sends and retrieves context instance state

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	ctx := context.Background()

	log.SetGlobalLogger(logrus.NewLogger())

	logger := log.WithField("main_field", "example")
	logger.Info("main method.")

	ctx = logger.ToContext(ctx)
	Foo(ctx)

}

func Foo(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Infof("%s method.", "main")
}
```

Contributing
--------
Every help is always welcome. Feel free do throw us a pull request, we'll do our best to check it out as soon as possible. But before that, let us establish some guidelines:

1. This is an open source project so please do not add any proprietary code or infringe any copyright of any sort.
2. Avoid unnecessary dependencies or messing up go.mod file.
3. Be aware of golang coding style. Use a lint to help you out.
4. Add tests to cover your contribution.
5. Add [godoc](https://elliotchance.medium.com/godoc-tips-tricks-cda6571549b) to your code. 
6. Use meaningful [messages](https://medium.com/@menuka/writing-meaningful-git-commit-messages-a62756b65c81) to your commits.
7. Use [pull requests](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/about-pull-requests).
8. At last, but also important, be kind and polite with the community.

Any submitted issue which disrespect one or more guidelines above, will be discarded and closed.


<hr>

Released under the [MIT License](LICENSE).
