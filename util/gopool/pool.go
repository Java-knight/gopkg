package gopool

import (
	"context"
	"sync"
	"sync/atomic"
)

/*
这里将pool 中需要执行的task使用一个双头单向链表维护了，将task作为一种资源，使用完后可以进行半初始化保存（指分配了内存空间，没有属性赋值）
Pool 增加了阈值机制，可以根据机器 CPU 动态扩展。Pool 是一个抽象的接口，它可以赋予多个实体
*/

type Pool interface {
	// Name 返回相应 pool 的名字
	Name() string

	// SetCapacity 设置 pool 的 goroutine 容量
	SetCapacity(cap int32)

	// Go 执行 f
	Go(f func())

	// CtxGo 执行 f, 并且接受 Context
	CtxGo(ctx context.Context, f func())

	// SetPanicHandler 设置 panic 处理器
	SetPanicHandler(handler func(context.Context, interface{}))

	// WorkerCount 返回正在运行的 worker 数量
	WorkerCount() int32
}

var taskPool sync.Pool

func init() {
	taskPool.New = newTask
}

// task 链表的一个 node。但是使用了 task 小写，封装性
type task struct {
	ctx context.Context
	f   func()

	next *task
}

// clean 清空 task 属性
func (t *task) clean() {
	t.ctx = nil
	t.f = nil
	t.next = nil
}

// Recycle 回收 task 资源（将其属性清空，然后放入taskPool【sync.pool】）
func (t *task) Recycle() {
	t.clean()
	taskPool.Put(t)
}

// newTask 创建一个task（分配了内存空间，没有属性赋值）
func newTask() interface{} {
	return &task{}
}

// taskList 双头链表（目前没有使用，但是后续考虑将task队列的特性移单独维护，二次开发可以增加新功能）
type taskList struct {
	sync.Mutex
	taskHead *task
	taskTail *task
}

type pool struct {
	name     string  // 名字
	capacity int32   // 容量
	config   *Config // 配置信息

	// task 链表（并发安全）
	taskHead  *task
	taskTail  *task
	taskLock  sync.Mutex
	taskCount int32 // taskCount 就是任务队列的任务数量（未开始的任务数量）

	// worker 数量（正在工作的goroutine）
	workerCount int32

	// panic 处理器
	panicHandler func(context.Context, interface{})
}

// NewPool 使用给定名称、上限和配置创建 pool
func NewPool(name string, cap int32, config *Config) Pool {
	p := &pool{
		name:     name,
		capacity: cap,
		config:   config,
	}
	return p
}

func (p *pool) Name() string {
	return p.name
}

func (p *pool) SetCapacity(cap int32) {
	atomic.StoreInt32(&p.capacity, cap)
}

func (p *pool) Go(f func()) {
	p.CtxGo(context.Background(), f)
}

func (p *pool) CtxGo(ctx context.Context, f func()) {
	t := taskPool.Get().(*task)
	t.ctx = ctx // 调用方传递的 context
	t.f = f     // 调用方需要执行的任务
	p.taskLock.Lock()
	if p.taskHead == nil { // 任务队列没有任务（链表）
		p.taskHead = t
		p.taskTail = t
	} else {
		p.taskTail.next = t
		p.taskTail = t
	}
	p.taskLock.Unlock()
	p.incrTaskCount() // task + 1

	// task开始运行的条件：（可以控制流量的大小，大流量就需要让task进行排队，机器内存升高，请求响应变慢？解决方案： 加机器，现在都是云上部署，k8s自动扩容，升高gopool的capacity即可）
	// 1. taskCount >= 阈值 && workerCount < capacity
	// 2. workerCount == 0
	if (atomic.LoadInt32(&p.taskCount) >= p.config.ScaleThreshold && p.WorkerCount() < atomic.LoadInt32(&p.capacity)) || p.WorkerCount() == 0 {
		p.incrWorkerCount()
		w := workerPool.Get().(*worker)
		w.pool = p
		w.run()
	}

}

func (p *pool) SetPanicHandler(handler func(context.Context, interface{})) {
	p.panicHandler = handler
}

func (p *pool) WorkerCount() int32 {
	return atomic.LoadInt32(&p.workerCount)
}

// worker 原子加1
func (p *pool) incrWorkerCount() {
	atomic.AddInt32(&p.workerCount, 1)
}

// worker 原子减1
func (p *pool) decrWorkerCount() {
	atomic.AddInt32(&p.workerCount, -1)
}

// task 原子加1
func (p *pool) incrTaskCount() {
	atomic.AddInt32(&p.taskCount, 1)
}

// task 原子减1
func (p *pool) decrTaskCount() {
	atomic.AddInt32(&p.taskCount, -1)
}
