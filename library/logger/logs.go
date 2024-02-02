package logger

import (
	"context"
)

var l *fileLogger

func New(loglevel string, skip int) {
	l = &fileLogger{
		zapLogger: initLogger(loglevel),
		skip:      skip,
	}
}

func WithContext(ctx context.Context) Logger {
	return &fileLogger{
		zapLogger: l.zapLogger,
		ctx:       ctx,
		skip:      l.skip,
	}
}

func WithCallerSkip(skip int) Logger {
	return &fileLogger{
		zapLogger: l.zapLogger,
		ctx:       l.ctx,
		skip:      skip,
	}
}

func Debug(message ...interface{}) {
	l.Debug(message...)
}

func Info(message ...interface{}) {
	l.Info(message...)
}

func Warn(message ...interface{}) {
	l.Warn(message...)
}

func Error(message ...interface{}) {
	l.Error(message...)
}

func DebugF(format string, message ...interface{}) {
	l.DebugF(format, message...)
}

func InfoF(format string, message ...interface{}) {
	l.InfoF(format, message...)
}

func WarnF(format string, message ...interface{}) {
	l.WarnF(format, message...)
}

func ErrorF(format string, message ...interface{}) {
	l.ErrorF(format, message...)
}
