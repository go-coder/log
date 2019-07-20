# 日志丢失问题

日志丢失有可能的原因，及其解决方案

> 1. 在goroutine中打印日志，主线程结束退出程序时，日志还没打印完。

这个问题需要用户自行维护程序的完整性，明确哪些程序是必须在主线程结束前退出的。建议阅读 [`Never start a goroutine without knowing how it will stop | Dave Cheney`](https://dave.cheney.net/2016/12/22/never-start-a-goroutine-without-knowing-how-it-will-stop)

> 2. 日志的后端实现使用了goroutine，主线程退出时，日志还没打印完。

这个问题是日志库本身需要解决的，在尚未得出一个完美的方案之前，我们先切换回加锁的实现方案。虽然这种方式效率低，但保证程序正确性应该是第一位的。
