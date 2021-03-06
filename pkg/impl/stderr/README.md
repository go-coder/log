# logr 后端设计

logr 在设计的时候充分考虑了扩展性，可以添加多种不同的后端方案，方便日志打印和日志归集。只要实现了 [`types.go`](https://github.com/go-coder/log/pkg/api/types.go) 中定义的`EntryWriter`接口，就可以直接切换后端，而不必担心日志前端的兼容问题。

这里，我们默认提供了 `stderr` 的[后端实现](https://github.com/go-coder/log/pkg/impl/stderr/stderr.go)。

## stderr 解决竞争的方案

log 前端部分采用只读类型解决了竞争问题，后端的竞争问题由后端自行解决。这里的解决思路同样是两个：只读、加锁（包括 chan 也是用到了锁）。

我们注意到，可能存在的竞争发生在调用 os.Stderr.WriteString 的时候，由于这不是一个原子操作，也就存在着一条日志没打完，另一条日志混入的情况。只读的解决办法看样子不行，我们这里考虑加互斥锁、或者使用 channel，下面比较一下两种方案的优势和不足：

使用 channel 的好处在于直观，所有的 log 请求被组织成了一个队列的形式，所有的日志汇集到了一个 goroutine 中进行处理；可能存在的问题是，当日志请求过多的时候，这个负责处理日志的 goroutine 可能会成为性能瓶颈。

加锁的话，日志请求在原来的 goroutine 中完成，避免了单个 goroutine 的性能问题，但是会影响当前 goroutine 的性能。

我们假定日志组织良好，不会达到 goroutine 性能瓶颈。这样，采用 channel 的方式排队处理请求，把日志处理从普通 goroutine 中提取出来，由专门的 goroutine 来处理，以期提高性能。

### 思考1

如果使用的是无缓冲chan的话，日志库性能并不高，因为阻塞队列必须每产生一条日志就打印下来，这其实是一个同步操作。带缓冲的chan能够提高性能。（这里其实是假设所有的 goroutine 是在一个线程内部了，被分配到一个核上执行）

ps: goroutine 是否能够充分利用多核CPU，把打印日志的 goroutine 单独放到另一个核中？这样的话就不存在阻塞问题了，性能会很好。

### 思考2

单独启动一个 goroutine 负责日志后端打印的话，存在一个问题：当主程序执行完一条打印日志命令后很快结束程序的时候，所有 goroutine 都会被杀死，但是这时候可能日志还没打印完毕，会有丢失日志的情况出现。我们需要一个同步的操作来等待日志打印完毕再结束进程。
