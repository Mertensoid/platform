package logging

// Importance levels int defer
type LogLevel int

// Importance levels constants
const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

// Logger interface. Different methods for each importance level
type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})

	Debug(string)
	Debugf(string, ...interface{})

	Info(string)
	Infof(string, ...interface{})

	Warn(string)
	Warnf(string, ...interface{})

	Panic(string)
	Panicf(string, ...interface{})
}
