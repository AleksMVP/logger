package logger

import (
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(writer io.Writer) *Logger {
	baseLogger := logrus.New()
	lgr := &Logger{baseLogger}
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "2006-01-02 15:04:05"
	Formatter.FullTimestamp = true
	lgr.SetFormatter(Formatter)
	lgr.SetOutput(writer)
	return lgr
}

func getCallerInfo() (fn string, pkg string) {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown", "unknown"
	}

	name := runtime.FuncForPC(pc).Name()
	slash := strings.LastIndexByte(name, '/')
	if slash < 0 {
		slash = 0
	}
	dot := strings.LastIndexByte(name[slash:], '.') + slash

	return name[dot+1:], name[:dot]
}

func (l *Logger) Infof(format string, args ...interface{}) {
	fn, pkg := getCallerInfo()
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": fn,
	}).Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	fn, pkg := getCallerInfo()
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": fn,
	}).Warnf(format, args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	fn, pkg := getCallerInfo()
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": fn,
	}).Warningf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	fn, pkg := getCallerInfo()
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": fn,
	}).Errorf(format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	fn, pkg := getCallerInfo()
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": fn,
	}).Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	fn, pkg := getCallerInfo()
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": fn,
	}).Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	fn, pkg := getCallerInfo()
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": fn,
	}).Error(args...)
}

func (l *Logger) LogInfo(pkg string, method string, msg string) {
	l.WithFields(logrus.Fields{
		"source":  pkg,
		"function": method,
	}).Info(msg)
}

func (l *Logger) LogWarn(pkg string, method string, msg string) {
	l.WithFields(logrus.Fields{
		"package":  pkg,
		"function": method,
	}).Warn(msg)
}

// Deprecated: use LogWarn
func (l *Logger) LogWarning(pkg string, method string, msg string) {
	l.LogWarn(pkg, method, msg)
}

func (l *Logger) LogError(pkg string, method string, err error) {
	l.WithFields(logrus.Fields{
		"package":  pkg,
		"function": method,
	}).Error(err)
}

func (l *Logger) LogAccess(r *http.Request, status int, worktime time.Duration) {
	l.WithFields(logrus.Fields{
		"method":      r.Method,
		"status":      status,
		"remote_addr": r.RemoteAddr,
		"work_time":   worktime.String(),
	}).Info(r.URL.Path)
}