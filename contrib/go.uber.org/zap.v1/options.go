package zap

type Options struct {
	Console struct {
		// enable/disable console logging
		Enabled bool
		// console log level
		Level string
		// console formatter TEXT/JSON
		Formatter string
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
		// file formatter TEXT/JSON
		Formatter string
	}
}
