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
	logs []Log
}
