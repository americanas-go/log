
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

#### Printf
logs a message at level Info (Logrus and Zap) and Debug (Zerolog) on the standard logger.

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

#### Tracef
logs a message at level Trace (Logrus and Zerolog) and Debug (Zap) on the standard logger.

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

#### Trace
logs a message at level Trace (Logrus and Zerolog) and Debug (Zap) on the standard logger.

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

#### Debugf
logs a message at level Debug on the standard logger.

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

#### Debug
logs a message at level Debug on the standard logger.

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

#### Infof
logs a message at level Info on the standard logger.

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

#### Info
logs a message at level Info on the standard logger.

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

#### Warnf
logs a message at level Warn on the standard logger.

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

#### Warn
logs a message at level Warn on the standard logger.

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

#### Errorf
logs a message at level Error on the standard logger.

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

#### Error
logs a message at level Error on the standard logger.

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

#### Fatalf
logs a message at level Fatal on the standard logger, then calls os.Exit(1).

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

#### Fatal
logs a message at level Fatal on the standard logger, then calls os.Exit(1).

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

#### Panicf
logs a message at level Panic on the standard logger.

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

#### Panic
logs a message at level Panic on the standard logger.

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

#### WithFields
creates an entry from the standard logger and adds multiple fields to it.

```go
package main

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	
	logrus.NewLogger()
	
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
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	logrus.NewLogger()
	log.WithField("hello", "world").
		Info("main method.")
}
```

#### WithTypeOf
creates an entry from the standard logger and adds type and package information to it.

```go
package main

import (
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
	"github.com/americanas-go/log"
)

func main() {
	logrus.NewLogger()
	s := S{}
	s.Foo()
}

type S struct {}

func (s *S) Foo() {
	logger := log.WithTypeOf(s)
	logger.Info("main method.")
}
```

#### ToContext
sends the state of the instance to the context.

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	ctx := context.Background()

	logrus.NewLogger()

	logger := log.WithField("main_field", "example")
	logger.Info("main method.")

	ctx = logger.ToContext(ctx)

}

func Foo(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Infof("%s method.", "main")
}
```

#### FromContext
returns a Logger from context.

```go
package main

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/sirupsen/logrus.v1"
)

func main() {
	ctx := context.Background()

	logrus.NewLogger()
	
	logger := log.WithField("main_field", "example")
	logger.Info("main method.")

	ctx = logger.ToContext(ctx)
	Foo(ctx)
}

func Foo(ctx context.Context) {
	logger := log.FromContext(ctx)
	logger.Infof("%s method.", "foo")
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