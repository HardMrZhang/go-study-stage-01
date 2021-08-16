package day08

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func hello() {
	fmt.Println("hello goroutine")
}

//启动单个goroutine
func Test_goroutine_01(t *testing.T) {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("hello main")
	time.Sleep(time.Second)
}

//启动多个goroutine
var wg sync.WaitGroup

func hello2(i int) {
	defer wg.Done() //goroutine结束就登记
	fmt.Println("hello goroutine", i)
}
func Test_goroutine_02(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello2(i)
	}
	wg.Wait()
}

//runtime.Gosched()
func Test_goroutine_03(t *testing.T) {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func Test_goroutine_04(t *testing.T) {
	//go func() {
	//	defer fmt.Println("A defer")
	//	func() {
	//		defer fmt.Println("B.defer")
	//		//结束协程
	//		runtime.Goexit()
	//		defer fmt.Println("C.defer")
	//		fmt.Println("B")
	//	}()
	//	fmt.Println("A")
	//
	//}()
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for i := 0; i < 2; i++ {
		fmt.Println("hello")
	}

}

//runtime.GOMAXPROCS
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
}
func Test_goroutine_05(t *testing.T) {
	//两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务
	//runtime.GOMAXPROCS(1)
	//go a()
	//go b()
	//time.Sleep(time.Second)

	//将逻辑核心数设为2，此时两个任务并行执行
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func Test_goroutine_06(t *testing.T) {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}

func Test_goroutine_07(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 10
	fmt.Println("发送成功")
}

func Test_goroutine_08(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}

/**
优雅的从通道循环取值
*/
func Test_goroutine_09(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}

/**
Goroutine池
*/

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

//创建工作池
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			//执行运算
			for job := range jobChan {
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func Test_goroutine_10(t *testing.T) {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	for i := 0; i < 10; i++ {
		id++
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}

}

/**
定时器
*/
func Test_goroutine_11(t *testing.T) {
	//1.timer基本使用
	//timer1 := time.NewTimer(2 * time.Second)
	//t1 := time.Now()
	//fmt.Printf("t1:%v\n", t1)
	//t2 := <-timer1.C
	//fmt.Printf("t2:%v\n", t2)

	//验证timer只响应1次
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到")
	}

	//timer实现延时的功能
	//A
	//time.Sleep(time.Second)
	////B
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	////C
	//<-time.After(3 * time.Second)
	//fmt.Println("3秒到")

	//停止定时器
	//timer4 := time.NewTimer(2 * time.Second)
	//go func() {
	//	<-timer4.C
	//	fmt.Println("d定时器执行了")
	//}()
	//b := timer4.Stop()
	//if b {
	//	fmt.Println("timer4关闭", b)
	//}

	//重置定时器
	//timer5 := time.NewTimer(3 * time.Second)
	//timer5.Reset(1 * time.Second)
	//fmt.Println(time.Now())
	//fmt.Println(<-timer5.C)
	//for {
	//}

}

//Ticker：时间到了，多次执行
func Test_goroutine_12(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
			}
		}
	}()

	for {
	}
}

/**
select可以同时监听一个或多个channel，直到其中一个channel ready
*/
func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "test2"
}
func Test_goroutine_13(t *testing.T) {
	out1 := make(chan string)
	out2 := make(chan string)
	go test1(out1)
	go test2(out2)
	select {
	case s1 := <-out1:
		fmt.Println("s1:", s1)
	case s2 := <-out2:
		fmt.Println("s2:", s2)
	default:

	}
}

/**
如果多个channel同时ready，则随机选择一个执行
*/
func Test_goroutine_14(t *testing.T) {
	c := make(chan int, 1)
	c2 := make(chan string, 1)
	go func() {
		c2 <- "hello"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c <- 1
	}()

	select {
	case v := <-c:
		fmt.Println("int:", v)
	case v := <-c2:
		fmt.Println("string:", v)
	}
	fmt.Println("main 结束")
}

/**
判断管道有没有满
*/
func Test_goroutine_15(t *testing.T) {
	// 创建管道
	out1 := make(chan string, 2)
	// 子协程写数据
	go write(out1)
	//取数据
	for s := range out1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("writer hello")
		default:
			fmt.Println("channel full")
			//close(ch)

		}
		time.Sleep(time.Millisecond * 500)
	}
}

/**
使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，
等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的
*/
var x int64
var lock sync.Mutex

func add2() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func Test_goroutine_16(t *testing.T) {

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

/**
读写互斥锁
*/

var (
	y      int64
	wgg    sync.WaitGroup //同步等待组 sync.WaitGroup
	llock  sync.Mutex
	rwlock sync.RWMutex
)

func w() {
	rwlock.Lock() //加写锁
	fmt.Println("---------", y)
	y = y + 1
	fmt.Println("=========", y)
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()
	wgg.Done()
}
func read() {
	rwlock.RLock()               //加读锁
	time.Sleep(time.Millisecond) //假设读操作耗时1毫秒
	rwlock.RUnlock()
	wgg.Done()
}
func Test_goroutine_17(t *testing.T) {

	start := time.Now()
	for i := 0; i < 10; i++ {
		wgg.Add(1)
		go w()
	}
	for i := 0; i < 1000; i++ {
		wgg.Add(1)
		go read()
	}
	wgg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start), y)
}

/**

 */
var z int64
var l sync.Mutex
var wwgg sync.WaitGroup

// 普通版加函数
func add1() {
	// x = x + 1
	z++ // 等价于上面的操作
	wwgg.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	z++
	l.Unlock()
	wwgg.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&z, 1)
	wwgg.Done()
}

func Test_goroutine_18(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wwgg.Add(1)
		//go add1() // 普通版add函数 不是并发安全的
		//go mutexAdd() // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wwgg.Wait()
	end := time.Now()
	fmt.Println(z)
	fmt.Println(end.Sub(start))
}
