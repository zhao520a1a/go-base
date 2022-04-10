package xlog

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	// log count
	cnTrace int64
	cnDebug int64
	cnInfo  int64
	cnWarn  int64
	cnError int64
	cnFatal int64
	cnPanic int64
	// log count stat stamp
	cnStamp int64

	slogMutex sync.Mutex

	logs []string
)

func init() {
	atomic.StoreInt64(&cnStamp, time.Now().Unix())
}

func addLogs(log string) {
	slogMutex.Lock()
	defer slogMutex.Unlock()
	logs = append(logs, log)
	if len(logs) > 10 {
		logs = logs[len(logs)-10:]
	}
}

func getLogs() []string {

	slogMutex.Lock()
	defer slogMutex.Unlock()

	tmp := make([]string, len(logs))
	copy(tmp, logs)
	logs = []string{}
	return tmp
}

// LogStat TODO: 统计日志打印情况，后期都可以去掉
func LogStat() (map[string]int64, []string) {
	st := map[string]int64{
		"TRACE": atomic.SwapInt64(&cnTrace, 0),
		"DEBUG": atomic.SwapInt64(&cnDebug, 0),
		"INFO":  atomic.SwapInt64(&cnInfo, 0),
		"WARN":  atomic.SwapInt64(&cnWarn, 0),
		"ERROR": atomic.SwapInt64(&cnError, 0),
		"FATAL": atomic.SwapInt64(&cnFatal, 0),
		"PANIC": atomic.SwapInt64(&cnPanic, 0),

		"STAMP": atomic.SwapInt64(&cnStamp, time.Now().Unix()),
	}

	return st, getLogs()
}
