# 基本日志库

## 功能介绍

[`go-logr/logr`](https://github.com/go-logr/logr) 的一个基础实现方案。

## 设计思路

+ 前后端分离，前端实现 logr 接口，后端使用 EntryWriter 接口
+ 提供了 os.Stderr 输出的实现
+ 加入命令行参数来筛选 Info 级别
+ 打印日志必须的字段，包括：日志级别、时间日期、文件名、行号、prefix、message、key-value消息
+ 使用 key-value 形式输出日志，便于后端以 json 等格式自动化处理
+ 传入指针时，自动取出数据打印

## 可能的使用场景

+ 远程日志：pubsub、 gRPC、HTTP网页 等形式
+ AB测试：同时运行新旧版本的程序，自动对比日志不同，从中发现可能存在的问题
+ 日志归集：将多个不同程序的日志汇总到一个地址进行分析

## 前端的竞争问题

考虑两种使用场景：

场景一：不同级别的日志之间是否存在竞争

```golang
log.V(1).Info()
go func() {
    log.V(2).Info()
}()
```

场景二：同级别日志之间是否存在竞争

```golang
log.Info()
go func() {
    log.Info()
}()
```

1. logr.V / logr.WithName / logr.WithValues

    实现的时候，都是在副本中进行的操作，只需要考虑 copy 函数是否有竞争问题即可。

    copy 函数所作的工作只是简单把当前 log 的数据成员拷贝到另一个 log 对象中去。修改发生在拷贝而来的副本内，并且修改完之前，新的 log对象并不会开始承担打印任务，因此不存在竞争问题。

1. logr.Enabled

    判断的是当前日志级别是否在阈值之内，因此只要当前日志级别不会被另一个 goroutine 修改就好。因为日志级别的设置由 logr.V 来完成，logr.V 每次修改都是生成新的副本，不会改变当前 log 对象的级别，所以不存在竞争。

1. logr.Info / logr.Error

    每次调用时，都是生成一个 Entry，由后端的 EntryWriter 实际承担日志打印任务。生成 Entry 的时候，由于所有的数据都不会被修改，所以不存在竞争；后端 EntryWriter 自行保证自己不存在竞争。

### 结论

logr 通过 log 对象只读的方式避免了竞争问题。所有对 log 进行改变的操作都是在新生成的副本之上进行的，从而保证了不会发生竞争。

## 后端竞争问题

由后端自行控制，stderr 的竞争问题分析参见 [README.md](https://github.com/go-coder/log/pkg/impl/stderr/README.md)
