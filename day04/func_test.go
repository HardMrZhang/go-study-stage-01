package main

import (
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func func1() {
	fmt.Println("这是函数func1")
}
func Test_func_001(t *testing.T) {
	/**
	匿名:没有名字
		匿名函数:没有名字的函数
		匿名变量:没有名字的变量

		定义一个匿名函数,直接进行调用。通常只能使用一次,也可以将匿名函数赋值给某个函数变量,那么可以调用多次.

	匿名函数:
		1.将函数作为另一个函数的参数(回调函数)
		2.将函数作为另一个函数的返回值
	*/
	//调用函数
	func1()
	func2 := func1
	//调用函数func2
	func2()

	//匿名函数
	func() {
		fmt.Println("我是一个匿名函数")
	}()
	func3 := func() { fmt.Println("我也是一个匿名函数") }
	func3()

	//带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//带返回值的匿名函数
	res := func(a, b int) int {
		return a + b
	}(10, 29)
	fmt.Println("res:", res)

}

func visit(l []int, f func(int)) { //使用visist将整个遍历过程进行封装,想要获取遍历期间的切片值,只需要向visit函数传入一个回调参数

	for _, v := range l {
		f(v)
	}
}

//使用匿名函数作为回调函数
func Test_func_01(t *testing.T) {

	var list = []int{1, 2, 3, 4, 6}
	visit(list, func(i int) {
		fmt.Println("遍list结果为:", i)
	})
}

//使用匿名函数实现操作封装
func Test_func_02(t *testing.T) {
	var skillParam = flag.String("eat", "eat", "skill perform")

	var skill = map[string]func(){
		"eat": func() {
			fmt.Println("eat true")
		},
		"read": func() {
			fmt.Println("read true")
		},
		"sleep": func() {
			fmt.Println("sleep true")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

}

//结构体实现接口
type Invoker interface {
	read(interface{})
}
type Stu struct { //定义一个空的stu结构体,没有任何属性,主要为了展示实现invoker中的read方法
}

func (s *Stu) read(v interface{}) { //read为结构体的方法,打印传入interface{}类型的值
	fmt.Println("from struct", v)
}

func Test_func_03(t *testing.T) {
	//声明接口变量
	var invoker Invoker
	//实例化结构体
	//s := new(Stu)
	s := &Stu{}
	//将实例化的结构体赋值到接口
	invoker = s
	//使用接口调用stu结构体中的read方法
	invoker.read("hello")
}

//函数实现接口

//函数定义为类型
type FuncReader func(interface{})

//实现Invoker中的Read
func (f FuncReader) read(v interface{}) {
	f(v) //FuncReader中的reader()方法被调用与func(interface{})无关,还需手动调用
}
func Test_func_04(t *testing.T) {
	//声明接口变量
	var invoker Invoker
	//将匿名函数转为FuncReader类型,再赋值给接口

	invoker = FuncReader(func(v interface{}) {
		fmt.Println("from func", v)
	})
	invoker.read("hello")
}

func Test_func_05(t *testing.T) {
	/**
	高阶函数:
		根据go语言数据类型特点,可以将一个函数func1作为另一个函数func2的参数
		那么func2为高阶函数,func1为回调函数
	*/
	fmt.Printf("%T\n", add) //func(int,int)int
	fmt.Printf("%T\n", op)  //func(int,int,func(int,int) int) int
	r1 := add(1, 2)
	fmt.Println("add：", r1)
	r2 := op(1, 2, add)
	fmt.Println("op：", r2)

	f1 := func(a, b int) int {
		return a * b
	}
	r3 := op(2, 5, f1)
	fmt.Println("使用匿名函数：", r3)

	r5 := op(10, 0, func(i int, i2 int) int {
		if i2 == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return i / i2
	})
	fmt.Println("r5:", r5)

}
func add(a, b int) int {
	return a + b
}

func op(a, b int, fun func(int, int) int) int {
	//fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}

//闭包
func Test_func_06(t *testing.T) {
	/**
	在闭包内部修改引用的变量
	*/
	str := "hello world"
	f := func() {
		str = "cz" //在匿名函数中并没有定义str,str定义是在匿名函数声明之前,此时str就被引用到匿名函数中形成闭包
	}
	f2 := func() {
		str = "zzy"
	}
	f()
	f2()
	fmt.Println(str)
}
func Test_func_07(t *testing.T) {
	/**
	闭包的记忆效应
		被捕获到的变量让闭包自身拥有记忆效应,闭包中的逻辑可以修改闭包捕获的变量,变量会跟随闭包的声明周期一直存在,
		闭包就如同变量一样拥有了记忆效应
	*/

	culRes := cul(1)
	fmt.Println(culRes())
	fmt.Printf("%p\n", &culRes)

	culRes2 := cul(10)
	fmt.Println(culRes2())
	fmt.Printf("%p\n", &culRes2)

	culRes3 := cul(100)
	fmt.Println(culRes3())
	fmt.Printf("%p\n", &culRes3)
}

func cul(a int) func() int {
	return func() int {
		a++
		return a
	}
}

//函数的可变参数
func Test_func_08(t *testing.T) {
	f(1, 2, 3, 4)
	f(2, 3, 4)
	var a = map[string]int{"one": 1, "two": 2}
	f2(a)
	//在可变参数函数中传递参数
	print(a)
}

func f(b int, args ...int) {
	for k, v := range args {
		fmt.Println(k, v)
	}
}
func f2(args ...interface{}) {
	for k, v := range args {
		fmt.Println(k, v)
	}
}

func print(args ...interface{}) {
	f2(args...)
}

//defer
func Test_func_09(t *testing.T) {
	defer readFile()
	fmt.Println("1111")
}

func readFile() bool {
	f, err := os.Open("/Users/zhangchaoyin/IdeaProjects/go-study-stage-01/day03/aa.txt")
	if err != nil {
		return false
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return false
	}
	fmt.Println(info.Size())
	return true
}

//延迟方法
type Student struct {
	Name    string
	Context string
}

func (s *Student) reduce() {
	fmt.Printf("%s %s \n", s.Name, s.Context)
}
func Test_func_10(t *testing.T) {
	s := Student{
		Name:    "zzy",
		Context: "golang",
	}
	defer s.reduce()
	fmt.Println("这是cz")
}

//延迟参数

func printA(a int) {
	fmt.Println("printA value：", a)
}
func Test_func_11(t *testing.T) {

	a := 5
	defer printA(a)
	a = 10
	fmt.Println(a)

}

/**
defer测试题
*/

func Test_func_12(t *testing.T) {

	e1()
	e2()
	e3()
}
func e1() {
	var err error //nil
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	errors.New("e3 defer err")
}

//计算函数执行时间
func Test_func_13(t *testing.T) {
	start := time.Now()
	sum := 0
	for i := 0; i < 10000; i++ {
		sum++
	}
	end := time.Now().Sub(start)
	//end := time.Since(start)

	fmt.Println("耗时为:", end)

}

//sha1 md5
func Test_func_14(t *testing.T) {

	str := "hello world!"

	md5String := md5.New()
	md5String.Write([]byte(str))
	res := md5String.Sum([]byte(""))
	fmt.Printf("%x\n\n", res)

	sha1String := sha1.New()
	sha1String.Write([]byte(str))
	res1 := sha1String.Sum([]byte(""))
	fmt.Printf("%x\n\n", res1)

}

//递归函数
func Test_func_15(t *testing.T) {
	fmt.Println("result:", factorial(7))
}
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
