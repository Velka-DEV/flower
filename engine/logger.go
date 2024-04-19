package engine

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

type Log struct {
	Level       LogLevel `json:"level"`
	Message     string   `json:"message"`
	Time        string   `json:"time"`
	ExecutionID string   `json:"execution_id"`
}

type Logger interface {
	Log(level LogLevel, message string, args ...interface{})
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

type DefaultLogger struct {
	logs        []Log
	executionId string
}

func NewDefaultLogger(executionId string) *DefaultLogger {
	return &DefaultLogger{
		logs:        []Log{},
		executionId: executionId,
	}
}

func (l *DefaultLogger) Log(level LogLevel, message string, args ...interface{}) {
	l.logs = append(l.logs, Log{
		Level:       level,
		Message:     message,
		ExecutionID: l.executionId,
	})
}

func (l *DefaultLogger) Debug(message string, args ...interface{}) {
	l.Log(LogLevelDebug, message, args...)
}

func (l *DefaultLogger) Info(message string, args ...interface{}) {
	l.Log(LogLevelInfo, message, args...)
}

func (l *DefaultLogger) Warn(message string, args ...interface{}) {
	l.Log(LogLevelWarn, message, args...)
}

func (l *DefaultLogger) Error(message string, args ...interface{}) {
	l.Log(LogLevelError, message, args...)
}
