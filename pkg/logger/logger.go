package logger

import glog "log"

var log Logger = &baseLog{}

// Logger 最简结构的 logger, 借用这个实现简单记录
type Logger interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// SetLogger 设置 log
func SetLogger(logger Logger) {
	log = logger
}

func GetLogger() Logger {
	return log
}

var (
	Infof  = log.Infof
	Debugf = log.Debugf
	Warnf  = log.Warnf
	Errorf = log.Errorf
	Fatalf = log.Fatalf
)

// baseLog 默认实现，使用原始的log，可替换
type baseLog struct {
}

func (b baseLog) Infof(format string, args ...interface{}) {
	glog.Printf(format, args...)
}

func (b baseLog) Debugf(format string, args ...interface{}) {
	glog.Printf(format, args...)
}

func (b baseLog) Warnf(format string, args ...interface{}) {
	glog.Printf(format, args...)
}

func (b baseLog) Errorf(format string, args ...interface{}) {
	glog.Printf(format, args...)
}

func (b baseLog) Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}
