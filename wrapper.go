package log

import (
	"sync"
)

// A global variable so that l functions can be directly accessed.
var (
	l  Logger = Noop{}
	mu sync.RWMutex
)

// NewLogger returns an instance of logger.
// Deprecated: prefer SetGlobalLogger
func NewLogger(logger Logger) {
	mu.Lock()
	defer mu.Unlock()
	l = logger
}

func SetGlobalLogger(logger Logger) {
	mu.Lock()
	defer mu.Unlock()
	l = logger
}

// Printf logs a templated message.
//
// For templating details see implementation doc.
func Printf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Printf(format, args...)
}

// Tracef logs a templated message at trace level.
//
// For templating details see implementation doc.
func Tracef(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Tracef(format, args...)
}

// Trace logs a message at trace level.
func Trace(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Trace(args...)
}

// Debugf logs a templated message at debug level.
//
// For templating details see implementation doc.
func Debugf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Debugf(format, args...)
}

// Debug logs a message at debug level.
func Debug(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Debug(args...)
}

// Infof logs a templated message at info level.
//
// For templating details see implementation doc.
func Infof(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Infof(format, args...)
}

// Info logs a message at info level.
func Info(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Info(args...)
}

// Warnf logs a templated message at warn level.
//
// For templating details see implementation doc.
func Warnf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Warnf(format, args...)
}

// Warn logs a message at warn level.
func Warn(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Warn(args...)
}

// Errorf logs a templated message at error level.
//
// For templating details see implementation doc.
func Errorf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Errorf(format, args...)
}

// Error logs a message at error level.
func Error(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Error(args...)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Panicf(format, args...)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Panic(args...)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Fatal(args...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()

	l.Fatalf(format, args...)
}

// WithField adds a key and value to logger.
func WithField(key string, value interface{}) Logger {
	mu.RLock()
	defer mu.RUnlock()

	return l.WithField(key, value)
}

// WithError adds an error as a field to logger
func WithError(err error) Logger {
	mu.RLock()
	defer mu.RUnlock()

	return l.WithError(err)
}

// WithFields adds fields to logger.
func WithFields(keyValues Fields) Logger {
	mu.RLock()
	defer mu.RUnlock()

	return l.WithFields(keyValues)
}

// WithTypeOf adds type information to logger.
func WithTypeOf(obj interface{}) Logger {
	mu.RLock()
	defer mu.RUnlock()

	return l.WithTypeOf(obj)
}

// GetLogger returns instance of Logger.
func GetLogger() Logger {
	mu.RLock()
	defer mu.RUnlock()

	return l
}
