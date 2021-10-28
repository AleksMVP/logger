package logger

import (
	"net/http"
	"time"
)

type ILogger interface {
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	LogInfo(pkg string, method string, msg string)
	LogWarn(pkg string, method string, msg string)
	LogError(pkg string, method string, err error)

	// Deprecated: don't use this
	LogAccess(r *http.Request, status int, worktime time.Duration) 
	// Deprecated: use LogWarn
	LogWarning(pkg string, method string, msg string)
}