package day09

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

type iface struct {
	tab  *itab
	data unsafe.Pointer
}
type itab struct {
	inter uintptr
	_type uintptr
	link  uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

type eface struct {
	_type uintptr
	data  unsafe.Pointer
}

func Test_reflect_01(t *testing.T) {
	var r io.Reader
	fmt.Printf("initial r: %T, %v\n", r, r)

	tty, _ := os.OpenFile("/Users/zhangchaoyin/IdeaProjects/go-study-stage-01/day09/reflect.md", os.O_RDWR, 0)
	fmt.Printf("tty: %T, %v\n", tty, tty)

	// 给 r 赋值
	r = tty
	fmt.Printf("r: %T, %v\n", r, r)

	rIface := (*iface)(unsafe.Pointer(&r))
	fmt.Printf("r: iface.tab._type = %#x, iface.data = %#x\n", rIface.tab._type, rIface.data)

	// 给 w 赋值
	var w io.Writer
	w = r.(io.Writer)
	fmt.Printf("w: %T, %v\n", w, w)

	wIface := (*iface)(unsafe.Pointer(&w))
	fmt.Printf("w: iface.tab._type = %#x, iface.data = %#x\n", wIface.tab._type, wIface.data)

	// 给 empty 赋值
	var empty interface{}
	empty = w
	fmt.Printf("empty: %T, %v\n", empty, empty)

	emptyEface := (*eface)(unsafe.Pointer(&empty))
	fmt.Printf("empty: eface._type = %#x, eface.data = %#x\n", emptyEface._type, emptyEface.data)

}

//1.提供一个结构体
type Person struct {
	Name string
	Age  int
	Sex  string
}

//2.提供一个方法
func (p Person) Say(msg string) {
	fmt.Println("Hello..", msg)
}

func (p Person) PrintInfo() {
	fmt.Println("姓名：", p.Name, "年龄：", p.Age, "性别：", p.Sex)
}
func Test_reflect_02(t *testing.T) {
	p1 := Person{"王二狗", 30, "男"}
	//反射使用 TypeOf 和 ValueOf 函数从接口中获取目标对象信息
	//1.获取对象的类型
	t1 := reflect.TypeOf(p1)
	fmt.Println(t1)                   //day09.Person
	fmt.Println("p1的类型是：", t1.Name()) //调用t.Name方法来获取这个类型的名称
	k1 := t1.Kind()                   //struct
	fmt.Println(k1)
	//2.获取值，如果是结构体类型，获取的是字段的值
	v1 := reflect.ValueOf(p1) //{王二狗 30 男}
	fmt.Println(v1)
	if t1.Kind() == reflect.Struct {
		//是结构体类型，获取里面的字段名字
		fmt.Println(t1.NumField()) //3
		for i := 0; i < t1.NumField(); i++ {
			field := t1.Field(i)
			//fmt.Println(field) //{Name  string  0 [0] false},{Age  int  16 [1] false},{Sex  string  24 [2] false}
			val := v1.Field(i).Interface() //通过interface方法来取出这个字段所对应的值
			fmt.Printf("字段名字：%s,字段类型：%s,字段数值：%v\n", field.Name, field.Type, val)
		}
	}

	//2.操作方法
	for i := 0; i < t1.NumMethod(); i++ {
		m := t1.Method(i)
		fmt.Println(m.Name, m.Type) //Hello func(day09.Person)
		/*
		   {Hello  func(day09.Person) <func(day09.Person) Value> 0}
		   {PrintInfo  func(day09.Person) <func(day09.Person) Value> 1}
		*/
	}

	m1 := v1.MethodByName("Say")
	args := []reflect.Value{reflect.ValueOf("干啥呢？")}
	m1.Call(args)

	m2 := v1.MethodByName("PrintInfo")
	m2.Call(nil)
}

func Test_reflect_03(t *testing.T) {
	var x float64 = 1.0
	//v := reflect.ValueOf(x)
	//v.SetFloat(2.0)
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(2.0)
	fmt.Println(v)
	fmt.Println(x)
}
