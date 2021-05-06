package zerolog

type Options struct {
	// formatter TEXT/JSON
	Formatter string
	Console   struct {
		// enable/disable console logging
		Enabled bool
		// console log level
		Level string
	}
	File struct {
		Enabled bool
		// enable/disable file logging
		Level string
		// log path
		Path string
		// log filename
		Name string
		// log file max size (MB)
		MaxSize int
		// enabled/disable file compress
		Compress bool
		// file max age
		MaxAge int
	}
}
