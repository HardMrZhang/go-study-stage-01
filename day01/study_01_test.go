package day01

import (
	"fmt"
	"math"
	"math/cmplx"
	"testing"
)

const NAME = "yyasyfa"

//变量声明
func Test_01(t *testing.T) {
	a := 1
	fmt.Println(&a)
	a, b := 20, "aa"
	fmt.Println(&a, a)
	fmt.Println(b)

	//var a = "aaaaaaaa"
	//fmt.Printf("%T ",unsafe.Sizeof(a))
}

func Test_02(t *testing.T) {
	const Width = 10
	const Height = 20
	var area = Width * Height
	const a, b, c = 1, false, "str"
	fmt.Printf("面积为 %d", area)

}

func Test_03(t *testing.T) {
	//const (
	//	a = iota
	//	b
	//	c = "a"
	//	d
	//	e = iota
	//)
	//fmt.Println(a, b, c, d, e)
	const (
		a = iota //0
		b        //1
		c = "aa" //aa
		d        //aa
		e = iota //2
		_
		f //3
	)
	fmt.Println(a, b, c, d, e, f)
}

func Test_04(t *testing.T) {
	//按位或运算
	fmt.Printf("12 | 10 的十进制结果是%d, 二进制结果是%b\n", 12|10, 12|10)

	//按位异或运算
	fmt.Printf("12 ^ 10 的十进制结果是%d, 二进制结果是%b\n", 12^10, 12^10)
	//按位与运算
	fmt.Printf("12 & 10 的十进制结果是%d, 二进制结果是%b\n", 12&10, 12&10)

	//位移 左移运算
	fmt.Println(12 << 1)
	fmt.Println(12 << 2)
	fmt.Println(12 << 3)
	fmt.Println(12 << 4)
	//右移运算
	fmt.Println(12 >> 1)
	fmt.Println(12 >> 2)
	fmt.Println(12 >> 3)
}

func Test_05(t *testing.T) {
	type ByteSize float64
	const (
		_           = iota // 通过赋值给空白标识符来忽略值
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println(KB,
		MB)
}

func Test_06(t *testing.T) {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x, y)
	fmt.Println(real(x), real(y))
	fmt.Println(imag(x), imag(y))
}

func Test(t *testing.T) {
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}
