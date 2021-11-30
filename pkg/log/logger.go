package log

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

type Field interface{}

func init() {
	l, _ := zap.NewProduction()
	logger = l.Sugar()
}

type Logger interface {
	With(fields ...Field) Logger
	Debugf(fmt string, args ...interface{})
	Infof(fmt string, args ...interface{})
	Warningf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Panicf(fmt string, args ...interface{})
}

type logWrapper struct {
	log *zap.SugaredLogger
}

func (l logWrapper) With(fields ...Field) Logger {
	return logWrapper{l.log.With(fields)}
}

func (l logWrapper) Panicf(fmt string, args ...interface{}) {
	l.log.Panicf(fmt, args...)
}

func (l logWrapper) Debugf(fmt string, args ...interface{}) {
	l.log.Debugf(fmt, args...)
}

func (l logWrapper) Infof(fmt string, args ...interface{}) {
	l.log.Infof(fmt, args...)
}

func (l logWrapper) Warningf(fmt string, args ...interface{}) {
	l.log.Warnf(fmt, args...)
}

func (l logWrapper) Errorf(fmt string, args ...interface{}) {
	l.log.Errorf(fmt, args...)
}

func With(fields ...Field) Logger {
	arr := make([]interface{}, len(fields))
	for _, val := range fields {
		arr = append(arr, val)
	}
	return logWrapper{logger.With(arr...)}
}

func Debugf(fmt string, args ...interface{}) {
	logger.Debugf(fmt, args)
}

func Infof(fmt string, args ...interface{}) {
	logger.Infof(fmt, args)
}

func Warningf(fmt string, args ...interface{}) {
	logger.Warnf(fmt, args)
}

func Errorf(fmt string, args ...interface{}) {
	logger.Errorf(fmt, args)
}

func Panicf(fmt string, args ...interface{}) {
	logger.Panicf(fmt, args)
}
