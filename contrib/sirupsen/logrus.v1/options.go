package logrus

type Options struct {
	Formatter string // formatter TEXT/JSON
	Time      struct {
		Format string
	}
	Console struct {
		Enabled bool   // enable/disable console logging
		Level   string // console log level
	}
	File struct {
		Enabled  bool   // enable/disable file logging
		Level    string // log level
		Path     string // log path
		Name     string // log filename
		MaxSize  int    // log file max size (MB)
		Compress bool   // enabled/disable file compress
		MaxAge   int    // file max age
	}
}
