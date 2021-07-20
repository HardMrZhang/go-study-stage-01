package day01

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
	fmt.Println(NAME)
}
