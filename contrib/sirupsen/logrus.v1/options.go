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
	Redis struct {
		Enabled bool   // enable/disable redis logging
		Key     string // redis key
		Format  string // redis format
		Host    string // redis host
		App     string // redis app
		Port    int    // redis port
		DB      int    // redis db
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
