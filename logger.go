package logger

import (
	"net/http"
	"time"
)

type ILogger interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})

	LogInfo(pkg string, method string, msg string)
	LogWarn(pkg string, method string, msg string)
	LogWarning(pkg string, method string, msg string)
	LogError(pkg string, method string, err error)
	LogAccess(r *http.Request, status int, worktime time.Duration) 
}