// Package logr provides Logur integration for Logr interface.
package logr

import (
	"fmt"

	"github.com/go-logr/logr"
	"logur.dev/logur"

	"logur.dev/integration/logr/internal/keyvals"
)

// Logger is a Logr logger.
type Logger struct {
	logger       logur.Logger
	levelEnabler logur.LevelEnabler
	name         string
}

// New returns a new Logr logger.
func New(logger logur.Logger) logr.Logger {
	l := &Logger{
		logger: logger,
	}

	if levelEnabler, ok := logger.(logur.LevelEnabler); ok {
		l.levelEnabler = levelEnabler
	}

	return logr.New(l)
}

// Info logs a non-error message with the given key/value pairs as context.
func (l *Logger) Info(level int, msg string, keysAndValues ...interface{}) {
	if len(keysAndValues) > 0 {
		l.logger.Info(msg, keyvals.ToMap(keysAndValues))

		return
	}

	l.logger.Info(msg)
}

// Enabled tests whether this InfoLogger is enabled.
func (l *Logger) Enabled(level int) bool {
	if l.levelEnabler == nil {
		return true
	}

	return l.levelEnabler.LevelEnabled(logur.Info)
}

// Error logs an error, with the given message and key/value pairs as context.
func (l *Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	if len(keysAndValues) > 0 {
		l.logger.Error(msg, keyvals.ToMap(keysAndValues))

		return
	}

	l.logger.Error(msg)
}

// WithValues adds some key-value pairs of context to a logger.
func (l *Logger) WithValues(keysAndValues ...interface{}) logr.LogSink {
	if len(keysAndValues) == 0 {
		return l
	}

	return &Logger{
		logger:       logur.WithFields(l.logger, keyvals.ToMap(keysAndValues)),
		levelEnabler: l.levelEnabler,
		name:         l.name,
	}
}

// WithName adds a new element to the logger's name.
func (l *Logger) WithName(name string) logr.LogSink {
	return &Logger{
		logger:       l.logger,
		levelEnabler: l.levelEnabler,
		name:         fmt.Sprintf("%s-%s", l.name, name),
	}
}

func (l *Logger) Init(info logr.RuntimeInfo) {
}
