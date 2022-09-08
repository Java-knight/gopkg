package gopool

import (
	"fmt"
	"gopkg/util/logger"
	"runtime/debug"
	"sync"
)

var workerPool sync.Pool

func init() {
	workerPool.New = newWorker
}

type worker struct {
	pool *pool
}

func newWorker() interface{} {
	return &worker{}
}

func (w *worker) run() {
	go func() {
		for {
			var curTask *task
			w.pool.taskLock.Lock()
			if w.pool.taskHead != nil { // 链表头不为空
				curTask = w.pool.taskHead
				w.pool.taskHead = w.pool.taskHead.next
				w.pool.decrTaskCount()
			}
			if curTask == nil { // 链表头为空
				// 如果没有task，就退出(将资源回收了)
				w.close()
				w.pool.taskLock.Unlock()
				w.Recycle()
				return
			}
			w.pool.taskLock.Unlock()

			// 拦截器处理Panic
			func() {
				defer func() {
					if r := recover(); r != nil {
						if w.pool.panicHandler != nil {
							w.pool.panicHandler(curTask.ctx, r)
						} else {
							msg := fmt.Sprintf("GOPOOL: panic in pool: %s: %v: %s", w.pool.name, r, debug.Stack())
							logger.CtxErrorf(curTask.ctx, msg)
						}
					}
				}()
				curTask.f()
			}()

			curTask.Recycle()
		}
	}()
}

// 关闭一个worker
func (w *worker) close() {
	w.pool.decrWorkerCount()
}

// 清空整个worker 的 pool
func (w *worker) clean() {
	w.pool = nil
}

// Recycle 回收 worker
func (w *worker) Recycle() {
	w.clean()
	workerPool.Put(w)
}
