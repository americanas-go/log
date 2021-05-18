package zerolog

type Options struct {
	Formatter string // formatter TEXT/JSON
	Console   struct {
		Enabled bool   // enable/disable console logging
		Level   string // console log level
	}
	File struct {
		Enabled  bool   // enable/disable file logging
		Level    string // file log level
		Path     string // file log path
		Name     string // file log filename
		MaxSize  int    // file log file max size (MB)
		Compress bool   // enabled/disable file compress
		MaxAge   int    // file max age
	}
}

type Option func(options *Options)

func WithFormatter(value string) Option {
	return func(options *Options) {
		options.Formatter = value
	}
}

func WithConsoleEnabled(value bool) Option {
	return func(options *Options) {
		options.Console.Enabled = value
	}
}

func WithConsoleLevel(value string) Option {
	return func(options *Options) {
		options.Console.Level = value
	}
}

func WithFileEnabled(value bool) Option {
	return func(options *Options) {
		options.File.Enabled = value
	}
}

func WithFileLevel(value string) Option {
	return func(options *Options) {
		options.File.Level = value
	}
}

func WithFilePath(value string) Option {
	return func(options *Options) {
		options.File.Path = value
	}
}

func WithFileName(value string) Option {
	return func(options *Options) {
		options.File.Name = value
	}
}

func WithFileMaxSize(value int) Option {
	return func(options *Options) {
		options.File.MaxSize = value
	}
}

func WithFileCompress(value bool) Option {
	return func(options *Options) {
		options.File.Compress = value
	}
}

func WithFileMaxAge(value int) Option {
	return func(options *Options) {
		options.File.MaxAge = value
	}
}
