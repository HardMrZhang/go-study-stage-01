package day05

import (
	"fmt"
	"testing"
)

func Test_Point_01(t *testing.T) {

	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)

	fmt.Println(&b)
}

type name int8 //定义一个新的类型
type stu struct {
	a int
	b bool
	name
}

func Test_Point_02(t *testing.T) {
	a := new(stu)
	a.a = 1
	a.b = false
	a.name = 20
	fmt.Println(a.b, a.a, a.name)

	var b = stu{1, false, 2}
	var c *stu = &b
	fmt.Println(b.a, b.b, b.name, &b)
	fmt.Println(c.a, c.b, c.name, &c, (*c).a)
}

func Test_Point_03(t *testing.T) {
	a := 100
	b := &a
	fmt.Println(*b)
}

//双重指针
func Test_Point_04(t *testing.T) {
	var a int
	var ptr *int
	var pptr **int
	a = 3000

	//指针ptr地址
	ptr = &a

	pptr = &ptr

	fmt.Printf("a = %d \n", a)
	fmt.Printf("指针变量 *ptr = %v \n", ptr)
	fmt.Printf("指向指针变量的指针 *ptr = %v \n", pptr)

}

//操纵指针改变变量的值
func Test_Point_05(t *testing.T) {
	a := 100
	b := &a
	fmt.Printf("b-value: %v  b-pointer %p \n", *b, b)
	*b++
	fmt.Printf("b-value: %v  b-pointer %p ", *b, b)
}

//函数指针和指针函数
func Test_Point_06(t *testing.T) {
	//函数指针
	a := 100
	b := a
	b = 200
	fmt.Println(a, b)

	fmt.Printf("函数add的类型: %T\n", add)
	fmt.Println("打印add函数：", add(23, 1))

	var f func(int, int) int = add
	fmt.Println("f1的执行结果：", f(a, b))
	/*
	  结论：函数声明时，函数名本身就是一个指针。不需要加*。
	  函数与函数赋值时，属于浅拷贝，拷贝指针，两个指针同时指向同一块内存空间。
	*/

	//指针函数
	res1 := ptrFunc()
	fmt.Printf("res1:%p,%v\n", res1, *res1)
	res2 := res1
	fmt.Printf("res2:%p,%v\n", res2, *res2)

	res1[0] = 100
	fmt.Println(*res1)
	fmt.Println(*res2)
}

//定义一个函数
func add(a, b int) int {
	return a + b
}

//定义一个指针函数
func ptrFunc() *[4]int {
	var arr [4]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i*2 + 1
	}
	fmt.Printf("数组地址：%p , 数组内容：%v \n", &arr, arr)
	return &arr
}

func Test_Point_07(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)
	fmt.Printf("m的类型:%T 数据地址:%p m的地址%p", m, m, &m)
	m2 := m
	m2["a"] = 2
	p := &m
	fmt.Printf("%T\n", p)
	fmt.Printf("p1的地址是：%p,存储的内容是：%p\n", &p, p) //
	fmt.Println(*p)
	fmt.Println((*p)["name"])
}

func Test_Point_08(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)
	fmt.Printf("m的类型:%T 数据地址:%p m的地址%p", m, m, &m)
	m2 := m
	m2["a"] = 2
	p := &m
	fmt.Printf("%T\n", p)
	fmt.Printf("p1的地址是：%p,存储的内容是：%p\n", &p, p) //
	fmt.Println(*p)
	fmt.Println((*p)["name"])
}

func Test_Point_09(t *testing.T) {
	num := [5]int{1, 2, 3, 4, 5}
	//指针数组
	var p [5]*int
	for i := 0; i < len(num); i++ {
		temp := &num[i]
		fmt.Println(temp)
		p[i] = temp
	}
	fmt.Println(p)

	//数组指针
	var p2 *[5]int
	p2 = &num
	for k := range p2 {
		fmt.Println((*p2)[k])
	}
}

//指针传递参数
func change(a *int) {
	*a = 20
}
func Test_Point_10(t *testing.T) {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)

}

/**
尽量不要这么写
*/
func modify(arr *[3]int) {
	arr[0] = 90
}

func Test_Point_11(t *testing.T) {
	a := [3]int{89, 90, 91}
	modify(&a)
	fmt.Println(a)
}

//改用切片
func modify2(s []int) {
	s[0] = 900
}
func Test_Point_12(t *testing.T) {
	a := []int{89, 90, 91}
	modify2(a[:])
	fmt.Println(a)
}
