# gopool

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
>
问题2：为什么taskList和sync.pool都是存放task的容器集合，他们存放的是一份吗？为什么要存放在两个容器？
>
问题3：如何计算整个gopool的总任务量？
>
问题4：为什么task都去维护一个逻辑结构，而worker并没有呢？
>
 
## Related
[ants](https://github.com/panjf2000/ants) 