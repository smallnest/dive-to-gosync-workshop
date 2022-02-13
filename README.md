# 深入Go并发编程研讨课

企业级的并发库: [gofer](https://github.com/smallnest/gofer)

Go提供了我们便利的进行并发编程的工具、方法和同步原语，同时也提供给我们诸多的犯错的机会，也就是俗称的“坑”。即使是顶级Go开发的项目，比如Docker、Kubernetes、gRPC、etcd， 都是有经验丰富的Go开发专家所开发，也踩过不少的并发的坑，而且依然源源不断的继续踩着,即便是标准库也是这样。（参考 []()）

分析和总结并发编程中的陷阱，避免重复踩在别人的坑中，正式本次培训课的重要内容。只有深入了解并发原语的实现，全面了解它们的特性和限制场景，注意它们的局限和容易踩的坑，才能提高我们的并发编程的能力。通过了解和学习其他人的经验和贡献的项目和库，我们可以更好的扩展我们的视野，避免重复的造轮子，或者说我们可以制作更好的轮子。

语言的内存模型定义了对变量的读写的可见性，可以清晰而准确读写事件的`happen before`关系。对于我们，可以很好地分析和编排goroutine的运行，避免数据的竞争和不一致的问题。

通过本次课程，你可以:

- 了解基本同步原语的具体实现、hack同步原语进行扩展，了解它们的使用场景和坑，已经别人是怎么踩的
- 了解一些扩展的同步源于，对于标准库sync包的补充
- 对于规模很大的项目，分布式同步原语是必不可少的，带你了解便利的分布式同步原语
- atomic可以保证对数据操作的一致性，利用CAS可以设计lock-free的数据结构
- channel是Go语言进行并发编程的很好的工具，带你了解它的使用姿势
- 了解Go语言的内存模型


## 并发原语在Go中的应用综述

## 基本并发原语
- Mutex的实现、扩展功能和坑。
- RWMutex的实现、扩展功能和坑。
- Waitgroup的实现、坑
- Cond的使用和坑
- Once的实现和坑，单例的Eager/Lazy实现
- Pool的坑， net.Conn的池
- Map的实现、应用场景
- Context的一些问题

## 扩展并发原语
- 可重入锁
- 信号量
- SingleFlight及应用
- ErrGroup
- 自旋锁
- 文件锁
- 并发Map的多种实现

## 原子操作
- 原子操作的实现
- 操作的数据类型
- 提供的函数
- 通用Value类型
- 扩展的原子操作库

## 分布式并发原语
- 锁，Mutex, RWmutex实战
- 栅栏
- leader选举
- 队列
- STM
- 其它分布式并发库

## channel
- 常见易犯错的channel使用场景
- 三大使用场景
- Or-done模式的三种实现
- 扇入
- 扇出
- Tee
- Pipeline
- 流式处理

## Go内存模型
- init函数
- goroutine
- channel
- Mutex/RWMutex
- Waitgroup
- Once
- atomic

## 习题研讨

[Go concurrency quizzes](https://github.com/smallnest/go-concurrent-quiz)