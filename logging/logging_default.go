package logging

import (
	"fmt"
	"log"
)

// Default logger struct
type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

// Returns minimal importance level
func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

// Uses log.Logger method "Output"
func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		l.loggers[level].Output(2, message)
	}
}

// Trace mathod realize. Writes message with Trace importance level.
func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

// Tracef mathod realize. Writes formatted message with Trace importance level.
func (l *DefaultLogger) Tracef(template string, vals ...interface{}) {
	l.write(Trace, fmt.Sprintf(template, vals...))
}

// Debug mathod realize. Writes message with Debug importance level.
func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

// Debugf mathod realize. Writes formatted message with Debug importance level.
func (l *DefaultLogger) Debugf(template string, vals ...interface{}) {
	l.write(Debug, fmt.Sprintf(template, vals...))
}

// Info mathod realize. Writes message with Information importance level.
func (l *DefaultLogger) Info(msg string) {
	l.write(Information, msg)
}

// Infof mathod realize. Writes formatted message with Information importance level.
func (l *DefaultLogger) Infof(template string, vals ...interface{}) {
	l.write(Information, fmt.Sprintf(template, vals...))
}

// Warn mathod realize. Writes message with Warning importance level.
func (l *DefaultLogger) Warn(msg string) {
	l.write(Warning, msg)
}

// Warnf mathod realize. Writes formatted message with Warning importance level.
func (l *DefaultLogger) Warnf(template string, vals ...interface{}) {
	l.write(Warning, fmt.Sprintf(template, vals...))
}

// Panic mathod realize. Writes message with Fatal importance level.
func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

// Panicf mathod realize. Writes formatted message with Fatal importance level.
func (l *DefaultLogger) Panicf(template string, vals ...interface{}) {
	formattedMessage := fmt.Sprintf(template, vals...)
	l.write(Fatal, formattedMessage)
	if l.triggerPanic {
		panic(formattedMessage)
	}
}
