package logger

import (
	"net/http"
	"time"
)

type ILogger interface {
	LogInfo(pkg string, method string, msg string)
	LogWarning(pkg string, method string, msg string)
	LogError(pkg string, method string, err error)
	LogAccess(r *http.Request, status int, worktime time.Duration) 
}