package gopool

import (
	"context"
	"fmt"
	"sync"
)

const (
	DEAULT_CAPACITY = 10000
)

// defaultPool 全局默认 pool
var defaultPool Pool

var poolMap sync.Map

func init() {
	defaultPool = NewPool("gopool.Default", DEAULT_CAPACITY, NewConfig())
}

// Go 是代替 go 关键字（启动协程），它可以恢复/拦截 Panic
//
// 使用姿势：
// gopool.Go(func(arg interface{}) {
//    ....
// })
func Go(f func()) {
	CtxGo(context.Background(), f)
}

// CtxGo 最优雅的使用姿势，可以传递 context
func CtxGo(ctx context.Context, f func()) {
	defaultPool.CtxGo(ctx, f)
}

// SetCapacity 不建议使用，该函数会改变全局 pool 的容量，影响其他的调用者
func SetCapacity(cap int32) {
	defaultPool.SetCapacity(cap)
}

// SetPanicHandler 为全局 pool 设置 Panic handler
func SetPanicHandler(f func(context.Context, interface{})) {
	defaultPool.SetPanicHandler(f)
}

// WorkerCount 正在运行的 worker 数量
func WorkerCount() int32 {
	return defaultPool.WorkerCount()
}

// RegisterPool 向全局注册一个 newPool，GetPool 可以按照name获取已经注册的 pool。如果注册了相同 Name的 pool，则返回 error
func RegisterPool(p Pool) error {
	_, loaded := poolMap.LoadOrStore(p.Name(), p)
	if loaded {
		return fmt.Errorf("name: %s already registered", p.Name())
	}
	return nil
}

// GetPool 按照名称获取注册的 pool。如果未注册返回 nil
func GetPool(name string) Pool {
	pool, ok := poolMap.Load(name)
	if !ok {
		return nil
	}
	return pool.(Pool)
}
