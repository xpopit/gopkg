package cls

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Log level constants
const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
)

var levelMap = map[string]int{
	"DEBUG": LevelDebug,
	"INFO":  LevelInfo,
	"WARN":  LevelWarn,
	"ERROR": LevelError,
}

func getLogLevel() int {
	lv := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	if val, ok := levelMap[lv]; ok {
		return val
	}
	return LevelDebug // default
}
func callerInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func NewLogger(logname, logFilePath string) (*XLogger, error) {
	var infoWriter, warnWriter, errorWriter, debugWriter io.Writer

	if logFilePath == "stdout" || logFilePath == "" {
		// ถ้าไม่มี logFilePath ให้ใช้ stdout/stderr เท่านั้น
		infoWriter = os.Stdout
		warnWriter = os.Stdout
		errorWriter = os.Stderr
		debugWriter = os.Stdout
	} else {
		// สร้าง log directory ถ้ายังไม่มี
		logDir := filepath.Dir(logFilePath)
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %v", err)
		}

		// เปิดไฟล์ log
		logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %v", err)
		}

		// สร้าง multi writer เพื่อเขียนทั้งไฟล์และ stdout/stderr
		infoWriter = io.MultiWriter(logFile, os.Stdout)
		warnWriter = io.MultiWriter(logFile, os.Stdout)
		errorWriter = io.MultiWriter(logFile, os.Stderr)
		debugWriter = io.MultiWriter(logFile, os.Stdout)
	}

	return &XLogger{
		logLevel: getLogLevel(),
		info:     log.New(infoWriter, fmt.Sprintf("[info] [%s] ", logname), log.Ldate|log.Ltime),
		warn:     log.New(warnWriter, fmt.Sprintf("[warn] [%s] ", logname), log.Ldate|log.Ltime),
		err:      log.New(errorWriter, fmt.Sprintf("[error] [%s] ", logname), log.Ldate|log.Ltime),
		debug:    log.New(debugWriter, fmt.Sprintf("[debug] [%s] ", logname), log.Ldate|log.Ltime),
	}, nil
}
func (l *XLogger) Info(format string, v ...interface{}) {
	if l.logLevel <= LevelInfo {
		msg := fmt.Sprintf(format, v...)
		l.info.Printf("%s: %s", callerInfo(2), msg)
	}
}

func (l *XLogger) Warn(format string, v ...interface{}) {
	if l.logLevel <= LevelWarn {
		msg := fmt.Sprintf(format, v...)
		l.warn.Printf("%s: %s", callerInfo(2), msg)
	}
}

func (l *XLogger) Error(format string, v ...interface{}) {
	if l.logLevel <= LevelError {
		msg := fmt.Sprintf(format, v...)
		l.err.Printf("%s: %s", callerInfo(2), msg)
	}
}

func (l *XLogger) Debug(format string, v ...interface{}) {
	if l.logLevel <= LevelDebug {
		msg := fmt.Sprintf(format, v...)
		l.debug.Printf("%s: %s", callerInfo(2), msg)
	}
}
