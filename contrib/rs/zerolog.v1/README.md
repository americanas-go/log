zerolog.v1
=======

The package accepts a default constructor:
```go
// default constructor
logger := zerolog.NewLogger()
```
Or a constructor with multiple parameters using optional pattern:
```go
// multiple optional parameters constructor
logger := zerolog.NewLogger(
	zerolog.WithFormatter("TEXT"),
	zerolog.WithConsoleEnabled(true),
	zerolog.WithFilePath("/tmp"),
	...
)
```

This is the list of all the configuration functions supported by package:

#### Formatter
WithFormatter sets output format of the logs. Using TEXT/JSON.
```go
// text formatter
logger := zerolog.NewLogger(zerolog.WithFormatter("TEXT"))

// json formatter
logger := zerolog.NewLogger(zerolog.WithFormatter("JSON"))
```

#### Level
WithLevel sets logging level to any of these options below on the standard logger.
```go
// log level DEBUG
logger := zerolog.NewLogger(zerolog.WithLevel("DEBUG"))

// log level WARN
logger := zerolog.NewLogger(zerolog.WithLevel("WARN"))

// log level FATAL
logger := zerolog.NewLogger(zerolog.WithLevel("FATAL"))

// log level ERROR
logger := zerolog.NewLogger(zerolog.WithLevel("ERROR"))

// log level TRACE
logger := zerolog.NewLogger(zerolog.WithLevel("TRACE"))

// log level INFO
logger := zerolog.NewLogger(zerolog.WithLevel("INFO"))
```

#### Console
WithConsoleEnabled sets whether the standard logger output will be in console. Accepts multi writing (console and file).
##### Enabled
```go
// console enable true
logger := zerolog.NewLogger(zerolog.WithConsoleEnabled(true))

// console enable false
logger := zerolog.NewLogger(zerolog.WithConsoleEnabled(false))
```

#### File
WithFileEnabled sets whether the standard logger output will be in file. Accepts multi writing (file and console).
##### Enabled
```go
// file enable true
logger := zerolog.NewLogger(zerolog.WithFileEnabled(true))

// file enable false
logger := zerolog.NewLogger(zerolog.WithFileEnabled(false))
```

##### Path
WithFilePath sets the path where the file will be saved.
```go
// file path
logger := zerolog.NewLogger(zerolog.WithFilePath("/tmp"))
```

##### Name
WithFileName sets the name of the file.
```go
// file name
logger := zerolog.NewLogger(zerolog.WithFileName("application.log"))
```

##### MaxSize
WithFileMaxSize sets the maximum size in megabytes of the log file. It defaults to 100 megabytes.
```go
// file max size
logger := zerolog.NewLogger(zerolog.WithFileMaxSize(100))
```

##### Compress
WithFileCompress sets whether the log files should be compressed.
```go
// file compress true
logger := zerolog.NewLogger(zerolog.WithFileCompress(true))

// file compress false
logger := zerolog.NewLogger(zerolog.WithFileCompress(false))
```

##### MaxAge
WithFileMaxAge sets the maximum number of days to retain old log files based on the timestamp encoded in their filename.
```go
// file max age
logger := zerolog.NewLogger(zerolog.WithFileMaxAge(10))
```