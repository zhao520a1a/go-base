package xlog

import "context"

var (
	// 应用日志
	appLogger *XLogger
	// 统计日志
	statLogger *XLogger
)

func init() {
	InitAppLog("", "", InfoLevel)
	InitStatLog("", "")
}

// Logger 注入其他基础库的日志句柄
type Logger struct {
}

func GetLogger() *Logger {
	return &Logger{}
}

func (m *Logger) Printf(format string, items ...interface{}) {
	Errorf(context.Background(), format, items...)
}

// Logger 注入其他基础库的日志句柄
type InfoLogger struct {
}

func GetInfoLogger() *InfoLogger {
	return &InfoLogger{}
}

func (m *InfoLogger) Printf(format string, items ...interface{}) {
	Infof(context.Background(), format, items...)
}

