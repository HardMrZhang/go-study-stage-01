# Go的并发编程

------

##### 并发介绍

###### 进程,线程和协程

1. 进程
   - 进程是一个程序在一个数据集中的一次动态执行过程，可以简单理解为“正在执行的程序”，它是CPU资源分配和调度的独立单位。
   - 进程一般由程序、数据集、进程控制块三部分组成。我们编写的程序用来描述进程要完成哪些功能以及如何完成；数据集则是程序在执行过程中所需要使用的资源；进程控制块用来记录进程的外部特征，描述进程的执行变化过程，系统可以利用它来控制和管理进程，它是系统感知进程存在的唯一标志。 
   - 进程的局限是创建、撤销和切换的开销比较大。
2. 线程
   - 线程是在进程之后发展出来的概念。 线程也叫轻量级进程，
   - 它是一个基本的CPU执行单元，也是程序执行过程中的最小单元，由线程ID、程序计数器、寄存器集合和堆栈共同组成。一个进程可以包含多个线程。
   - 线程的优点是减小了程序并发执行时的开销，提高了操作系统的并发性能，缺点是线程没有自己的系统资源，只拥有在运行时必不可少的资源，但同一进程的各线程可以共享进程所拥有的系统资源，如果把进程比作一个车间，那么线程就好比是车间里面的工人。不过对于某些独占性资源存在锁机制，处理不当可能会产生“死锁”。
3. 协程
   - 协程是一种用户态的轻量级线程，又称微线程，英文名Coroutine，协程的调度完全由用户控制。人们通常将协程和子程序（函数）比较着理解。
   - 子程序调用总是一个入口，一次返回，一旦退出即完成了子程序的执行。

###### 并发和并行

1. 并发

   - 多线程程序在一个核的cpu上运行,就是并发

   ![image-20210805164435192](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210805164435192.png)

2. 并行

   - 多线程程序在多个核的cpu上运行,就是并行

   ![image-20210805164515320](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210805164515320.png)

###### goroutine只是由官方实现的超级‘线程池’

- 创建一个 goroutine 的栈内存消耗为 2 KB，实际运行过程中，如果栈空间不够用，会自动进行扩容。创建一个 thread 则需要消耗 1 MB 栈内存
- 对于一个用 Go 构建的 HTTP Server 而言，对到来的每个请求，创建一个 goroutine 用来处理是非常轻松的一件事。而如果用一个使用线程作为并发原语的语言构建的服务，例如Java 来说，每个请求对应一个线程则太浪费资源了，很快就会出 内存溢出错误。

###### goroutine 奉行通过通信来共享内存，而不是共享内存来通信

##### Goroutine

Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。

###### 使用goroutine

Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。

一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。

###### 启动单个goroutine

启动goroutine的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个go关键字

###### 启动多个goroutine

##### runtime包

1. runtime.gosched()
   - 让出CPU时间片，重新等待安排任务
2. runtime.Goexit()
   - 退出当前协程
3. runtime.GOMAXPROCS
   - Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数
   - Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
4. Go语言中的操作系统线程和goroutine的关系：
   - 一个操作系统线程对应用户态多个goroutine。
   - go程序可以同时使用多个操作系统线程。
   - goroutine和OS线程是多对多的关系，即m:n

Channel

单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是CSP，提倡通过通信共享内存而不是通过共享内存而实现通信。

如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

Go 语言中的通道是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

1. channel类型

   - channel是一种引用类型,var 变量名 chan 元素类型

     ```go
     var 变量名 chan 元素类型
     var c1 chan int   // 声明一个传递整型的通道
     var c2 chan bool  // 声明一个传递布尔型的通道
     var c3 chan []int // 声明一个传递int切片的通道
     ```

2. 创建channel

   - 通道是引用类型,通道类型空值是nil

     ```go
     var ch chan int
     fmt.Println(ch) //<nil>
     ```

     声明的通道后需要使用make函数初始化之后才能使用

     ```go
     c1 := make(chan int)
     c2 := make(chan bool)
     c3 := make(chan []int)
     ```

3. channel操作

   - send

     ```go
     ch <- 10
     ```

   - receive

     ```go
     x := <-ch //从ch中接收值并赋给x
     <-ch 			//从ch中接收值,忽略结果		
     ```

   - Close

     ```go
     close(ch)
     ```

     关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的

     - 对一个关闭的通道再发送值就会panic
     - 对一个关闭的通道进行接收会一直获取值直到通道为空
     - 对一个关闭并且没有值的通道执行接收操作会得到对应类型的零值
     - 关闭一个已经关闭的通道会导致panic

4. 无缓冲的通道

   ![image-20210809143912242](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210809143912242.png)

5. 有缓冲的通道

   ![image-20210809145048849](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210809145048849.png)

6. 单向通道

   - 单向 channel 变量的声明非常简单，只能写入数据的通道类型为`chan<-`，只能读取数据的通道类型为`<-chan`

     ```go
     var 通道实例 chan<- 元素类型    // 只能写入数据的通道
     var 通道实例 <-chan 元素类型    // 只能读取数据的通道
     chan 在前面只能写入 chan 在后面只能读取
     ```

     

##### Goroutine池

- 本质上是一个生产者消费者模型
- 可以控制goroutine数量,防止暴涨

##### 定时器

##### select

###### select多路复用

在某些场景下我们需要同时从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。你也许会写出如下代码使用遍历的方式来实现:

```go
for{
    // 尝试从ch1接收值
    data, ok := <-ch1
    // 尝试从ch2接收值
    data, ok := <-ch2
    …
}
```

这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。

select的使用类似于switch语句，它有一系列case分支和一个默认的分支。每个case会对应一个通道的通信（接收或发送）过程。select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体格式如下

```go
 select {
    case <-chan1:
       // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
       // 如果成功向chan2写入数据，则进行该case处理语句
    default:
       // 如果上面都没有成功，则进入default处理流程
    }
```

- select可以同时监听一个或多个channel，直到其中一个channel ready
- 如果多个channel同时ready，则随机选择一个执行
- 可以用于判断管道是否存满

##### 并发安全和锁

有时候在Go代码中可能会存在多个goroutine同时操作一个资源（临界区），这种情况会发生竞态问题（数据竞态）

```go
var x int64
var wg sync.WaitGroup

func add() {
    for i := 0; i < 5000; i++ {
        x = x + 1
    }
    wg.Done()
}
func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}
```

上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符。

1. 互斥锁

   互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。

2. 读写互斥锁

   互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。

   读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

##### Sync

1.  sync.WaitGroup

   在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：

   | 方法名                          | 功能                |
   | ------------------------------- | ------------------- |
   | (wg * WaitGroup) Add(delta int) | 计数器+delta        |
   | (wg *WaitGroup) Done()          | 计数器-1            |
   | (wg *WaitGroup) Wait()          | 阻塞直到计数器变为0 |

   - sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成

   - sync.WaitGroup是一个结构体，传递的时候要传递指针。

##### 原子操作(atomic包)

1. 原子操作

   代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。Go语言中原子操作由内置的标准库sync/atomic提供。

2. atomic包

| **方法**                                                     | **解释**       |
| ------------------------------------------------------------ | -------------- |
| func LoadInt32(addr *int32) (val int32)<br/>func LoadInt64(addr `*int64`) (val int64)<br/>func LoadUint32(addr`*uint32`) (val uint32)<br/>func LoadUint64(addr`*uint64`) (val uint64)<br/>func LoadUintptr(addr`*uintptr`) (val uintptr)<br/>func LoadPointer(addr`*unsafe.Pointer`) (val unsafe.Pointer) | 读取操作       |
| func StoreInt32(addr `*int32`, val int32)<br/>func StoreInt64(addr `*int64`, val int64)<br/>func StoreUint32(addr `*uint32`, val uint32)<br/>func StoreUint64(addr `*uint64`, val uint64)<br/>func StoreUintptr(addr `*uintptr`, val uintptr)<br/>func StorePointer(addr `*unsafe.Pointer`, val unsafe.Pointer) | 写入操作       |
| func AddInt32(addr `*int32`, delta int32) (new int32)<br/>func AddInt64(addr `*int64`, delta int64) (new int64)<br/>func AddUint32(addr `*uint32`, delta uint32) (new uint32)<br/>func AddUint64(addr `*uint64`, delta uint64) (new uint64)<br/>func AddUintptr(addr `*uintptr`, delta uintptr) (new uintptr) | 修改操作       |
| func SwapInt32(addr `*int32`, new int32) (old int32)<br/>func SwapInt64(addr `*int64`, new int64) (old int64)<br/>func SwapUint32(addr `*uint32`, new uint32) (old uint32)<br/>func SwapUint64(addr `*uint64`, new uint64) (old uint64)<br/>func SwapUintptr(addr `*uintptr`, new uintptr) (old uintptr)<br/>func SwapPointer(addr `*unsafe.Pointer`, new unsafe.Pointer) (old unsafe.Pointer) | 交换操作       |
| func CompareAndSwapInt32(addr `*int32`, old, new int32) (swapped bool)<br/>func CompareAndSwapInt64(addr `*int64`, old, new int64) (swapped bool)<br/>func CompareAndSwapUint32(addr `*uint32`, old, new uint32) (swapped bool)<br/>func CompareAndSwapUint64(addr `*uint64`, old, new uint64) (swapped bool)<br/>func CompareAndSwapUintptr(addr `*uintptr`, old, new uintptr) (swapped bool)<br/>func CompareAndSwapPointer(addr `*unsafe.Pointer`, old, new unsafe.Pointer) (swapped bool) | 比较并交换操作 |
|                                                              |                |

