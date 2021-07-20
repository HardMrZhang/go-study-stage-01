package day02

import (
	"fmt"

	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a 的地址: %p ,长度 :%d %#v\n", a, len(a), a)
	b := a[2:3]
	b[0] = 1
	fmt.Printf("a 的地址: %p ,长度 :%d %#v\n", a, len(a), a)
	fmt.Printf("b 的地址: %p ,长度 :%d %#v\n", b, len(b), b)
	for i := 0; i < 2; i++ {
		b = append(b, 30)
		//fmt.Printf("a1 的地址: %p ,长度 :%d %#v\n", a, len(a), a)
		//fmt.Printf("b 的地址: %p ,长度 :%d %#v\n", b, len(b), b)
	}
}
func Test_Slice_01(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	a := arr[1:3]
	b := arr[:3]
	c := arr[1:]
	d := arr[:]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	a1 := make([]int, 2)
	b1 := make([]int, 2, 10)
	fmt.Println(a1, b1)
	fmt.Println(len(a1), len(b1))
}

//修改切片
func Test_Slice_02(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:5]
	fmt.Println("修改之前:", s)
	fmt.Println("修改之前:", arr)
	for k, _ := range s {
		s[k]++
	}
	fmt.Println("修改之后:", s)
	fmt.Println("修改之后:", arr)
}

//len和cap
func Test_Slice_03(t *testing.T) {
	var sarr = make([]int, 2, 10)
	fmt.Printf("len = %d , cap = %d , slice = %d \n", len(sarr), cap(sarr), sarr)

	//空切片在为初始化之前默认为nil
	var s []int
	fmt.Printf("len = %d , cap = %d , slice = %d \n", len(s), cap(s), s)
	if s == nil {
		fmt.Println("切片是空的")
	}
}

//append
func Test_Slice_04(t *testing.T) {
	var a []int
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a, &a[0])
	//在切片的开头添加元素
	//在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次，
	//因此，从切片的开头添加元素的性能要比从尾部追加元素的性能差很多。
	a = append([]int{0}, a...)
	fmt.Println(a, &a[1])

	//append的链式操作
	//var b []int
	a = append(a[:2], append([]int{28}, a[2:]...)...) //在第2个位置插入28
	fmt.Println(a)
	a = append(a[:2], append([]int{18, 19, 20}, a[2:]...)...)
	//每个添加操作中的第二个 append`调用都会创建一个临时切片，
	//并将 a[2:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:2] 中。
	fmt.Println(a)
}

//copy
func Test_Slice_05(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
	fmt.Println(slice2)

	/*
	   copy(),
	       切片是引用类型，传递的是地址

	   深拷贝和浅拷贝
	       深拷贝：拷贝的是数据
	           值类型都是深拷贝，基本类型,数组

	       浅拷贝：拷贝的是地址
	           引用类型默认都是浅拷贝，切片，map
	*/

	s3 := []int{1, 2, 3, 4}
	s4 := s3 //浅拷贝

	s5 := make([]int, 4, 4)
	copy(s5, s3) //深拷贝
	fmt.Println(s4)
	fmt.Println(s5)
	s3[0] = 100
	fmt.Println(s4)
	fmt.Println(s5)

}
