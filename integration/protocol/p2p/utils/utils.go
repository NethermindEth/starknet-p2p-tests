package utils

import (
	"fmt"
)

type Network struct {
	Name string `json:"name" validate:"required"`
}

const (
	Kilobyte = 1024
	Megabyte = 1024 * Kilobyte
	Gigabyte = 1024 * Megabyte
	Terabyte = 1024 * Gigabyte
)

type SimpleLogger interface {
	Debugw(msg string, keysAndValues ...any)
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, keysAndValues ...any)
	Errorw(msg string, keysAndValues ...any)
	Tracew(msg string, keysAndValues ...any)
}

type TestLogger func(format string, args ...interface{})

type TestSimpleLogger struct {
	Logger TestLogger
}

func (l *TestSimpleLogger) log(level, msg string, keysAndValues ...interface{}) {
	if len(keysAndValues) == 0 {
		l.Logger("%s: %s", level, msg)
		return
	}

	logMsg := fmt.Sprintf("%s: %s", level, msg)
	for i := 0; i < len(keysAndValues); i += 2 {
		key := fmt.Sprintf("%v", keysAndValues[i])
		var value interface{} = "MISSING"
		if i+1 < len(keysAndValues) {
			value = keysAndValues[i+1]
		}
		logMsg += fmt.Sprintf(" %s=%v", key, value)
	}

	l.Logger("%s", logMsg)
}

func (l *TestSimpleLogger) Logf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.Logger("INFO: %s", msg)
}

func (l *TestSimpleLogger) Debugw(msg string, keysAndValues ...any) {
	l.log("DEBUG", msg, keysAndValues...)
}

func (l *TestSimpleLogger) Infow(msg string, keysAndValues ...any) {
	l.log("INFO", msg, keysAndValues...)
}

func (l *TestSimpleLogger) Warnw(msg string, keysAndValues ...any) {
	l.log("WARN", msg, keysAndValues...)
}

func (l *TestSimpleLogger) Errorw(msg string, keysAndValues ...any) {
	l.log("ERROR", msg, keysAndValues...)
}

func (l *TestSimpleLogger) Tracew(msg string, keysAndValues ...any) {
	l.log("TRACE", msg, keysAndValues...)
}

var _ SimpleLogger = (*TestSimpleLogger)(nil)
