package logger

import (
	"context"
	"fmt"
	"log"
	"os"
)

const (
	STACK_NUM = 3 // 日志堆栈层数，比如出现panic, 日志会打印到第几层堆栈信息
)

var (
	_ Logger = (*localLogger)(nil)
)

// SetDefaultLogger 设置默认 logger。这不是并发安全的，那么它只能在 init 期间调用
func SetDefaultLogger(l Logger) {
	if l == nil {
		panic("logger must not be nil")
	}
	defaultLogger = l
}

// 使用 localLogger 包装 go原生的log。 os.Stderr 是 标准错误fd；log.LstdFlags 标准logger 初始值；log.Lshortfile 记录文件名和行号；log.Lmicroseconds 记录时间
var defaultLogger Logger = &localLogger{
	logger: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
}

type localLogger struct {
	logger *log.Logger
}

// lv: 调用方指定到日志等级；format: 日志格式；v: 日志信息（这里面可以规定调用几层堆栈信息 STACK_NUM）
func (ll *localLogger) logf(lv Level, format *string, v ...interface{}) {
	if level > lv {
		return
	}
	msg := lv.toString()
	if format != nil { // [Info]: logger【logger 中有%格式】
		msg += fmt.Sprintf(*format, v...)
	} else { // 【logger 中无%格式】
		msg += fmt.Sprint(v...)
	}

	ll.logger.Output(STACK_NUM, msg)
}

func (ll *localLogger) Trace(v ...interface{}) {
	ll.logf(LevelTrace, nil, v...)
}

func (ll *localLogger) Debug(v ...interface{}) {
	ll.logf(LevelDebug, nil, v...)
}

func (ll *localLogger) Info(v ...interface{}) {
	ll.logf(LeveInfo, nil, v...)
}

func (ll *localLogger) Notice(v ...interface{}) {
	ll.logf(LevelNotice, nil, v...)
}

func (ll *localLogger) Warn(v ...interface{}) {
	ll.logf(LevelWarn, nil, v...)
}

func (ll *localLogger) Error(v ...interface{}) {
	ll.logf(LevelError, nil, v...)
}

func (ll *localLogger) Fatal(v ...interface{}) {
	ll.logf(LevelFatal, nil, v...)
}

func (ll *localLogger) Tracef(format string, v ...interface{}) {
	ll.logf(LevelTrace, &format, v...)
}

func (ll *localLogger) Debugf(format string, v ...interface{}) {
	ll.logf(LevelDebug, &format, v...)
}

func (ll *localLogger) Infof(format string, v ...interface{}) {
	ll.logf(LeveInfo, &format, v...)
}

func (ll *localLogger) Noticef(format string, v ...interface{}) {
	ll.logf(LevelNotice, &format, v...)
}

func (ll *localLogger) Warnf(format string, v ...interface{}) {
	ll.logf(LevelWarn, &format, v...)
}

func (ll *localLogger) Errorf(format string, v ...interface{}) {
	ll.logf(LevelError, &format, v...)
}

func (ll *localLogger) Fatalf(format string, v ...interface{}) {
	ll.logf(LevelFatal, &format, v...)
}

func (ll *localLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelTrace, &format, v...)
}

func (ll *localLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelDebug, &format, v...)
}

func (ll *localLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LeveInfo, &format, v...)
}

func (ll *localLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelNotice, &format, v...)
}

func (ll *localLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelWarn, &format, v...)
}

func (ll *localLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelError, &format, v...)
}

func (ll *localLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	ll.logf(LevelFatal, &format, v...)
}
