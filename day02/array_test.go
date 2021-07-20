package day02

import (
	"fmt"
	"testing"
)

func Test_Arr_01(t *testing.T) {

	a := [3]int{1, 2, 3}
	b := [...]int{123}
	c := []int{1: 5, 6: 8}
	d := []int{1, 2, 8: 10}
	fmt.Printf("%v\n\n", a)
	fmt.Printf("%v\n\n", b)
	fmt.Printf("%v\n\n", c)
	fmt.Printf("%v\n\n", d)

}

/**
Go中的数组是值类型，而不是引用类型。这意味着当它们被分配给一个新变量时，
将把原始数组的副本分配给新变量。如果对新变量进行了更改，则不会在原始数组中反映。
*/

func Test_Arr_02(t *testing.T) {
	a := [...]string{"A", "B", "C", "D", "E"}
	b := a
	b[0] = "F"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	//当将数组传递给函数作为参数时，它们将通过值传递，而原始数组将保持不变
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before  ", num)
	changeLocal(num)
	fmt.Println("after ", num)
}
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("changeLocal func  ", num)

}

/**
冒泡排序
*/
func Test_Arr_03(t *testing.T) {
	arr := [5]int{5, 8, 1, 7, 10}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[i] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
