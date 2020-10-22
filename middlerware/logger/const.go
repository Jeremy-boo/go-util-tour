package logger

const (
	LogLevel   = "debug"
	FileName   = "./log/debug.log"
	MaxSize    = 2
	MaxAge     = 7
	MaxBackups = 10
)

// NewLogger used to initialize log file parameters
type NewLogger struct {
	Level       string
	FileName    string
	ServiceName string
	MaxSize     int
	MaxAge      int
	MaxBackups  int
}
