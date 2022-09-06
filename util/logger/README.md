

### 设计原则
使用了 localLogger 包装了 go 原生的 log库，提供了跟多级别的日志等级。
采用了面向对象的编程思想，基于接口编程，给外部暴露的是 Logger 接口，其实现是 localLogger 做的（小写显示了封装性）
```go
// 使用 localLogger 包装 go原生的log。 
// os.Stderr 是 标准错误fd；log.LstdFlags 标准logger 初始值；log.Lshortfile 记录文件名和行号；log.Lmicroseconds 记录时间
var defaultLogger Logger = &localLogger{
	logger: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
}
```

日志打印的核心方法：
```go
const (
    STACK_NUM = 3 // 日志堆栈层数，比如出现panic, 日志会打印到第几层堆栈信息
)

// lv: 调用方指定到日志等级；format: 日志格式；v: 日志信息（这里面可以规定调用几层堆栈信息 STACK_NUM）
func (ll *localLogger) logf(lv Level, format *string, v ...interface{}) {
	if level > lv {
		return
	}
	msg := lv.toString()
	if format != nil { // [Info]: logger【logger 中有%格式】
		msg += fmt.Sprintf(*format, v...)
	} else { // 【logger 中无%格式】
		msg += fmt.Sprintf(*format, v...)
	}
	// 输出，STACK_NUM 可以规定打印堆栈的层数
	ll.logger.Output(STACK_NUM, msg)
}
```

日志的使用
```go
// init 设置日志level
func init() {
	logger.SetLevel(logger.LevelTrace)
}

func main() {
	logger.Info("hahah")
	fmt.Println("husww")
}
```
生成的日志格式
```shell
2022/09/06 19:27:59.847775 logger.go:95: [Info][hahah]
```

### go语言注意事项
var 定义的变量在启动的时候就会初始化