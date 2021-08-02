package day06

import (
	"fmt"
	"testing"
)

type Info struct {
	Id      int
	content string
}
type Point struct {
	X    int
	Y    int
	Z    *Point
	Info *Info
}

func Test_struct_01(t *testing.T) {
	var p Point
	p.X = 10
	p.Y = 20
	//p1是指针 *Point
	p1 := new(Point)
	p1.Y = 200
	p1.X = 100

	p2 := &Point{}
	p2.X = 300
	p2.Y = 400
	fmt.Println(p, p1, p2)
}

func Test_struct_02(t *testing.T) {
	//键值对形式
	p := &Point{
		X: 0,
		Y: 0,
		Info: &Info{
			Id:      0,
			content: "",
		},
		Z: &Point{
			X:    0,
			Y:    0,
			Z:    nil,
			Info: nil,
		},
	}
	fmt.Println(p)

	//多个值形式
	info := &Info{1, "hello"}
	fmt.Println(info)
}

//匿名结构体
func Test_struct_03(t *testing.T) {
	msg := &struct { //定义部分
		id   int
		data string
	}{
		1, "hello",
	}
	printMsgType(msg)
}

func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

/**
go-模拟构造函数重载
*/

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{Name: name}
}
func NewCatByColor(color string) *Cat {
	return &Cat{Color: color}
}

/**
go-模拟父类子类构造并初始化
*/

type BuOuCat struct {
	Cat
}

//构造基类
func NewCat(name string) *Cat {
	return &Cat{Name: name}
}

//构造子类
func NewBuOuCat(color string) *BuOuCat {
	cat := &BuOuCat{}
	cat.Color = color
	return cat
}

//类型内嵌和结构体内嵌
type inners struct {
	a int
	b int
}

type outers struct {
	c int
	d float32
	int
	inners
}

func Test_struct_04(t *testing.T) {
	outer := new(outers)
	outer.c = 6
	outer.d = 7.5
	outer.int = 60
	outer.a = 5
	outer.b = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.d)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.a)
	fmt.Printf("outer.in2 is: %d\n", outer.b)
	// 使用结构体字面量
	outer2 := outers{6, 7.5, 60, inners{5, 10}}
	fmt.Println("outer2 is:", outer2)
}
