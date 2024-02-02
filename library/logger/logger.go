package logger

import "context"

type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Serve(...interface{})
	DebugF(string, ...interface{})
	InfoF(string, ...interface{})
	WarnF(string, ...interface{})
	ErrorF(string, ...interface{})
	WithContext(context.Context) Logger
	WithCallerSkip(int) Logger
}
