package day07

import (
	"fmt"
	"math"
	"testing"
)

//定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
	//CanWrite() bool
}

//定义文件结构,用于实现DataWriter
type file struct {
}

func (f *file) WriteData(data interface{}) error {
	//模拟写入数据
	fmt.Println("写入数据:", data)
	return nil
}

//func (f *file) CanWrite() bool {
//
//	return true
//}
func Test_Interface_01(t *testing.T) {
	f := new(file)
	//声明DataWriter接口
	var writer DataWriter
	// 将接口赋值f，也就是*file类型
	// *file 类型的f赋值给 DataWriter接口的writer,虽然两个变量类型不一致。
	//但是writer是一个接口，且f已经完全实现了DataWriter()的所有方法，因此赋值是成功的。
	writer = f
	writer.WriteData("hello")
}

/**
一个类型可以实现多个接口
*/

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}
type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}
func Test_Interface_02(t *testing.T) {
	var s Sayer
	var m Mover
	var a = dog{name: "旺财"}
	s = a
	m = a
	s.say()
	m.move()
}

/**
多个类型实现相同的接口
*/

type Drinker interface {
	drink()
}
type cat struct {
	name string
}
type human struct {
	name string
}

func (c *cat) drink() {
	fmt.Printf("%s会喝水\n", c.name)
}
func (h *human) drink() {
	fmt.Printf("%s会喝水\n", h.name)
}
func Test_Interface_03(t *testing.T) {
	var c = cat{name: "nico"}
	var h = human{name: "zzy"}
	var d Drinker
	d = &c
	d.drink()

	d = &h
	d.drink()
}

/**
接口的nil判断
*/

type MyImplement struct {
}

//实现fmt.Stringer的string方法
func (m *MyImplement) String() string {
	return "hi"
}

//在函数中返回fmt.Stringer接口
func GetStringer() fmt.Stringer {
	var s *MyImplement = nil
	if s == nil {
		return nil
	}
	return s
}

func Test_Interface_04(t *testing.T) {
	//判断返回值是否为nil
	fmt.Println(GetStringer())
	if GetStringer() == nil {
		fmt.Println("GetStringer() == nil")
	} else {
		fmt.Println("GetStringer() != nil")
	}
}

func Test_Interface_05(t *testing.T) {
	var x interface{}
	x = "aaa"
	//注意如果不接收第二个参数也就是上面代码中的ok，
	//断言失败时会直接造成一个panic。
	value, ok := x.(int)
	fmt.Println(value, ok)

}

/**
自定义错误类型
*/

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {
	return fmt.Sprintf("wrong!!,because \"%f\" is a negative number", e.Num)
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}
func Test_Interface_06(t *testing.T) {
	result, err := Sqrt(-31)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
