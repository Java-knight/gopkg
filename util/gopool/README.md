# gopool
![gopool](../../imgs/gopool_logo.png)
## Introduction
`gopool` 是一个高性能的 goroutine pool，目前是复用 goroutine 并限制 goroutine 的数量。

它是 `go` 关键字的替代品。

思考的第一个问题，作为高性能的 pool，肯定会有很大的并发量，所以需要考虑并发安全。
而且 高性能（使用合理的数据结构组织起来task的调度，减少开销）

## Features
- 高性能
- 自动恢复 Panic
- 限制 `goroutine` 数量
- 复用 `goroutine` 堆栈

## QuickStart

old:
```go
go func() {
	// do your job
}
```

new:
```go
gopool.Go func() {
	// do your job
}
```

## Design
整个gopool的运行原理：当一个task结束，worker会调用task.Recycle()方法，将task资源回收放入sync.pool中；
如果taskList中没有task（当head == tail时，还剩下一个没有任何赋值的task在taskList），
每个goroutine会将worker资源回收放入 sync.pool中

**[架构图](https://po-files.ks3-cn-beijing.ksyun.com/631a35e8f346fb55d8aadc94_e8874d5b1cf6.png)**
![架构图](https://po-files.ks3-cn-beijing.ksyun.com/631a35e8f346fb55d8aadc94_e8874d5b1cf6.png)
## Question
问题1：为什么这里的go func() {} 里面有一个loop，作用是什么？
> 不断的从 TaskList 中获取 Task进行执行。这个 worker 启动了，只要执行完当前 Task, 
> 就会继续获取新的 Task继续执行，知道 TaskList 中没有 Task了。
> worker 才会停止自己回收自己（清空属性，将其放入 sync.Pool 中）

问题2：为什么taskList和sync.pool都是存放task的容器集合，他们存放的是一份吗？为什么要存放在两个容器？
> TaskList 和 sync.pool 中存储的 Task 不是不同的。sync.pool 中是没有实际的 Task存储的，
> 即使有也是被重置属性后的 Task；而 TaskList 中是用户调用一次 `gopool.Go` 就会创建/从sync.pool获取 一个 Task，
> 将其连成一个链表（双头链表），使用 worker 去顺序消费 TaskList 中的 Task。每个 Task 被执行后会被重置放入 sync.pool

问题3：如何计算整个gopool的总任务量？
> TaskList 中表示已经申请还未开始做的任务；worker的数量是已经开始做的任务数量。
> 当前 `goppol` 的
> 
> 任务总量 = TaskCount + WorkerCount

问题4：为什么task都去维护一个逻辑结构，而worker并没有呢？
> 因为 Task 和 worker 的诉求不一样，Task 更想金主爸爸，我给你一个任务，
> 你尽快给我干完。那么为例公平起见就是排队，维护了一个 TaskList 尽心排队；
> 而 worker 更像是码头的工人，谁干完活了，就来领取下一个任务，谁先干完谁就来领。
> 活有大有小，但是worker 是没有只会的，只会干完就接下一个，所以也是公平的。
 
## Related
[ants](https://github.com/panjf2000/ants) 