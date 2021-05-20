zap.v1
=======

The package accepts a default constructor:
```go
// default constructor
logger := zap.NewLogger()
```
Or a constructor with multiple parameters using optional pattern:
```go
// multiple optional parameters constructor
logger := zap.NewLogger(
	zap.WithConsoleFormatter("TEXT"),
	zap.WithConsoleEnabled(true),
	zap.WithFilePath("/tmp"),
	...
)
```

This is the list of all the configuration functions supported by package:

#### Console
WithConsoleEnabled sets whether the standard logger output will be in console. Accepts multi writing (console and file).
##### Enabled
```go
// console enable true
logger := zap.NewLogger(zap.WithConsoleEnabled(true))

// console enable false
logger := zap.NewLogger(zap.WithConsoleEnabled(false))
```

#### Level
WithConsoleLevel sets console logging level to any of these options below on the standard logger.
```go
// log level DEBUG
logger := zap.NewLogger(zap.WithConsoleLevel("DEBUG"))

// log level WARN
logger := zap.NewLogger(zap.WithConsoleLevel("WARN"))

// log level FATAL
logger := zap.NewLogger(zap.WithConsoleLevel("FATAL"))

// log level ERROR
logger := zap.NewLogger(zap.WithConsoleLevel("ERROR"))

// log level TRACE
logger := zap.NewLogger(zap.WithConsoleLevel("TRACE"))

// log level INFO
logger := zap.NewLogger(zap.WithConsoleLevel("INFO"))
```

##### Formatter
WithConsoleFormatter sets output format of the console logs. Using TEXT/JSON.
```go
// text formatter
logger := zap.NewLogger(zap.WithConsoleFormatter("TEXT"))

// json formatter
logger := zap.NewLogger(zap.WithConsoleFormatter("JSON"))
```

#### File
WithFileEnabled sets whether the standard logger output will be in file. Accepts multi writing (file and console).
##### Enabled
```go
// file enable true
logger := zap.NewLogger(zap.WithFileEnabled(true))

// file enable false
logger := zap.NewLogger(zap.WithFileEnabled(false))
```

#### Level
WithFileLevel sets level logging to any of these options below on the standard logger.
```go
// log level DEBUG
logger := zap.NewLogger(zap.WithFileLevel("DEBUG"))

// log level WARN
logger := zap.NewLogger(zap.WithFileLevel("WARN"))

// log level FATAL
logger := zap.NewLogger(zap.WithFileLevel("FATAL"))

// log level ERROR
logger := zap.NewLogger(zap.WithFileLevel("ERROR"))

// log level TRACE
logger := zap.NewLogger(zap.WithFileLevel("TRACE"))

// log level INFO
logger := zap.NewLogger(zap.WithFileLevel("INFO"))
```

##### Path
WithFilePath sets the path where the file will be saved.
```go
// file path
logger := zap.NewLogger(zap.WithFilePath("/tmp"))
```

##### Name
WithFileName sets the name of the file.
```go
// file name
logger := zap.NewLogger(zap.WithFileName("application.log"))
```

##### MaxSize
WithFileMaxSize sets the maximum size in megabytes of the log file. It defaults to 100 megabytes.
```go
// file max size
logger := zap.NewLogger(zap.WithFileMaxSize(100))
```

##### Compress
WithFileCompress sets whether the log files should be compressed.
```go
// file compress true
logger := zap.NewLogger(zap.WithFileCompress(true))

// file compress false
logger := zap.NewLogger(zap.WithFileCompress(false))
```

##### MaxAge
WithFileMaxAge sets the maximum number of days to retain old log files based on the timestamp encoded in their filename.
```go
// file max age
logger := zap.NewLogger(zap.WithFileMaxAge(10))
```

##### Formatter
WithFileFormatter sets output format of the file logs. Using TEXT/JSON.
```go
// text formatter
logger := zap.NewLogger(zap.WithFileFormatter("TEXT"))

// json formatter
logger := zap.NewLogger(zap.WithFileFormatter("JSON"))
```

Example
--------

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
	//output: 2021-05-16T14:30:31.788-0300	info	runtime/proc.go:225	main method.	{"main_field": "example"}

	ctx = logger.ToContext(ctx)

	foo(ctx)

	withoutContext()
}

func foo(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("foo_field", "example")
	logger.Infof("%s method.", "foo")
	//output: 2021-05-16T14:30:31.788-0300	info	contrib/main.go:24	foo method.	{"main_field": "example", "foo_field": "example"}

	ctx = logger.ToContext(ctx)
	bar(ctx)
}

func bar(ctx context.Context) {
	logger := log.FromContext(ctx)

	logger = logger.WithField("bar_field", "example")

	logger.Infof("%s method.", "bar")
	//output: 2021-05-16T14:30:31.788-0300	info	contrib/main.go:37	bar method.	{"bar_field": "example", "main_field": "example", "foo_field": "example"}
}

func withoutContext() {
	log.Info("withoutContext method")
	//output: 2021-05-16T14:30:31.788-0300	info	contrib/main.go:50	withoutContext method
}
```