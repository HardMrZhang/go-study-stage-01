package main

import "fmt"

//结构体实现接口
type Invoker interface {
	read(interface{})
}
type Stu struct { //定义一个空的stu结构体,没有任何属性,主要为了展示实现invoker中的read方法
}

func (s *Stu) read(v interface{}) {
	fmt.Println("from struct", v)
}

func main() {
	//声明接口变量
	var invoker Invoker
	//实例化结构体
	s := new(Stu)
	//将实例化的结构体赋值到接口
	invoker = s
	invoker.read("hello")
}
