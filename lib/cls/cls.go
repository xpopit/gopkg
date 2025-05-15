package cls

import "log"

type CLS struct {
}
type Logger interface {
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
	Debug(format string, v ...interface{})
}

// สร้าง LineLogger struct
type XLogger struct {
	logLevel int
	info     *log.Logger
	warn     *log.Logger
	err      *log.Logger
	debug    *log.Logger
}
