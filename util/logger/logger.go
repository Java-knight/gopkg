package logger

import (
	"context"
	"fmt"
)

// Logger 是一个日志记录接口，提供带级别的记录功能
type Logger interface {
	Trace(v ...interface{})  // 追踪
	Debug(v ...interface{})  // 断点
	Info(v ...interface{})   // 信息
	Notice(v ...interface{}) // 注意
	Warn(v ...interface{})   // 警告
	Error(v ...interface{})  // 错误
	Fatal(v ...interface{})  // 失败

	// xxxf 支持 % 日志打印

	Tracef(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Noticef(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	// Ctxxxxf 传递 context 并支持 % 日志打印

	CtxTracef(ctx context.Context, format string, v ...interface{})
	CtxDebugf(ctx context.Context, format string, v ...interface{})
	CtxInfof(ctx context.Context, format string, v ...interface{})
	CtxNoticef(ctx context.Context, format string, v ...interface{})
	CtxWarnf(ctx context.Context, format string, v ...interface{})
	CtxErrorf(ctx context.Context, format string, v ...interface{})
	CtxFatalf(ctx context.Context, format string, v ...interface{})
}

// Level 定义日志消息的优先级。当 logger 是配置了一个级别时，任何具有较低日志级别（通过整数比较更小）的日志消息都不会输出。
type Level int

// 日志的级别
const (
	LevelTrace Level = iota
	LevelDebug
	LeveInfo
	LevelNotice
	LevelWarn
	LevelError
	LevelFatal
)

// SetLevel 设置日志的级别，低于该级别的日志将不会被输出。默认日志级别是 LevelTrace
func SetLevel(lv Level) {
	if lv < LevelTrace || lv > LevelFatal {
		panic("invalid level.")
	}
	level = lv
}

// Fatal 调用默认 logger 的 Fatal 方法，然后调用 os.Exit(1)
func Fatal(v ...interface{}) {
	defaultLogger.Fatal(v)
}

// Error 调用默认 logger 的 Error 方法
func Error(v ...interface{}) {
	if level > LevelError {
		return
	}
	defaultLogger.Error(v)
}

// Warn 调用默认 logger 的 Warn 方法
func Warn(v ...interface{}) {
	if level > LevelWarn {
		return
	}
	defaultLogger.Warn(v)
}

// Notice 调用默认 logger 的 Notice 方法
func Notice(v ...interface{}) {
	if level > LevelNotice {
		return
	}
	defaultLogger.Notice(v)
}

// Info 调用默认 logger 的 Info 方法
func Info(v ...interface{}) {
	if level > LeveInfo {
		return
	}
	defaultLogger.Info(v)
}

// Debug 调用默认 logger 的 Debug 方法
func Debug(v ...interface{}) {
	if level > LevelDebug {
		return
	}
	defaultLogger.Debug(v)
}

// Trace 调用默认 logger 的 Trace 方法
func Trace(v ...interface{}) {
	if level > LevelTrace {
		return
	}
	defaultLogger.Trace(v)
}

// Fatalf 调用默认 logger 的 Fatalf 方法，然后调用 os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

// Errorf 调用默认 logger 的 Errorf 方法
func Errorf(format string, v ...interface{}) {
	if level > LevelError {
		return
	}
	defaultLogger.Errorf(format, v)
}

// Warnf 调用默认 logger 的 Warnf 方法
func Warnf(format string, v ...interface{}) {
	if level > LevelWarn {
		return
	}
	defaultLogger.Warnf(format, v)
}

// Noticef 调用默认 logger 的 Noticef 方法
func Noticef(format string, v ...interface{}) {
	if level > LevelNotice {
		return
	}
	defaultLogger.Noticef(format, v)
}

// Infof 调用默认 logger 的 Infof 方法
func Infof(format string, v ...interface{}) {
	if level > LeveInfo {
		return
	}
	defaultLogger.Infof(format, v)
}

// Debugf 调用默认 logger 的 Debugf 方法
func Debugf(format string, v ...interface{}) {
	if level > LevelDebug {
		return
	}
	defaultLogger.Debugf(format, v)
}

// Tracef 调用默认 logger 的 Tracef 方法
func Tracef(format string, v ...interface{}) {
	if level > LevelTrace {
		return
	}
	defaultLogger.Tracef(format, v)
}

// CtxFatelf 调用默认 logger 的 CtxFatelf 方法，然后调用 os.Exit(1)
func CtxFatelf(ctx context.Context, format string, v ...interface{}) {
	defaultLogger.CtxFatalf(ctx, format, v...)
}

// CtxErrorf 调用默认 logger 的 CtxErrorf 方法
func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	if level > LevelError {
		return
	}
	defaultLogger.CtxErrorf(ctx, format, v...)
}

// CtxWarnf 调用默认 logger 的 CtxWarnf 方法
func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	if level > LevelWarn {
		return
	}
	defaultLogger.CtxWarnf(ctx, format, v...)
}

// CtxNoticef 调用默认 logger 的 CtxNoticef 方法
func CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	if level > LevelNotice {
		return
	}
	defaultLogger.CtxNoticef(ctx, format, v...)
}

// CtxInfof 调用默认 logger 的 CtxInfof 方法
func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	if level > LeveInfo {
		return
	}
	defaultLogger.CtxInfof(ctx, format, v...)
}

// CtxDebugf 调用默认 logger 的 CtxDebugf 方法
func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	if level > LevelDebug {
		return
	}
	defaultLogger.CtxDebugf(ctx, format, v...)
}

// CtxTracef 调用默认 logger 的 CtxTracef 方法
func CtxTracef(ctx context.Context, format string, v ...interface{}) {
	if level > LevelTrace {
		return
	}
	defaultLogger.CtxTracef(ctx, format, v...)
}

// level 对象
var level Level

var strs = []string{
	"[Trace]",
	"[Debug]",
	"[Info]",
	"[Notice]",
	"[Warn]",
	"[Error]",
	"[Fatal]",
}

// toString 打印日志等级，用途：每个 log 的开头[Trace]、[Debug]、[Info]...
func (lv Level) toString() string {
	if lv >= LevelTrace && lv <= LevelFatal {
		return strs[lv]
	}
	return fmt.Sprintf("[?%d]", lv)
}
