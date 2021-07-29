package day02

//go tool compile -N -l -S

import (
	"fmt"
	"testing"
	"time"
)

func Test_Circle_01(t *testing.T) {

	if num := 10; num%2 == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func Test_Switch_02(t *testing.T) {
	var grade int = 20
	var level string = "D"

	switch grade {
	case 90:
		level = "A"
	case 60:
		level = "B"
	case 30:
		level = "C"
	default:
		level = "D"
	}
	switch {
	case level == "A":
		fmt.Println("A")
	case level == "B":
		fmt.Println("B")
	case level == "C":
		fmt.Println("C")
	case level == "D":
		fmt.Println("D")
	}
}

func Test_Switch_03(t *testing.T) {

	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		fallthrough
	case 6:
		x += 20
		fmt.Println(x)
	}

}

/**
case中的表达式是可选的，可以省略。如果该表达式被省略，则被认为是switch true，
并且每个case表达式都被计算为true，并执行相应的代码块。
*/
func Test_Switch_04(t *testing.T) {
	num := 10
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("d")
	case num > 50 && num <= 100:
		fmt.Println("c")
	}
}

func Test_Switch_05(t *testing.T) {
	/*
	   switch的其他写法：
	   1.省略switch后的表达式，相当于直接作用在true
	   2.case后可以同时匹配多个数据,匹配上任意一个都可以执行该case分支
	   3.switch后支持多一条初始化语句
	*/

	//1.省略switch后的表达式
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}

	//2.case后有多个数值
	var str = "a"
	switch str {
	case "a", "b", "c", "d":
		fmt.Println("在其中")
	default:
		fmt.Println("不在其中")
	}

	//3.多一条初始化语句
	switch str1 := "golang"; str1 {
	case "java":
		fmt.Println("java")
	case "C++":
		fmt.Println("c++")
	default:
		fmt.Println("R")

	}
}

//判断type
func Test_Switch_06(t *testing.T) {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型 ：%T", i)
	case int:
		fmt.Printf("x的类型是int")
	case float64:
		fmt.Printf("x的类型是float64")
	case func(int):
		fmt.Printf("x是func(int)型")
	case bool, string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Println("未知类型")

	}
}

//select
func Test_Switch_07(t *testing.T) {
	ch := make(chan int)
	quit := make(chan bool)
	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}() //别忘了()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束")
}

func Test_For_08(t *testing.T) {
	var a = 5
	var b = 6
	nums := [10]int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("数组元素为 %d\n", nums[i])
	}
	for a > b {
		fmt.Println("a>b ")
	}
	//break
	for k, _ := range nums {
		if k > 3 {
			fmt.Println(nums[k])
			break
		}
	}
	//continue
	for _, v := range nums {
		if v%2 == 0 {
			fmt.Printf("能被二整除的是%d \n ", v)
			continue
		}
	}

}

//贴标签
func Test_For_09(t *testing.T) {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	var b = [5]int{1, 3, 5, 7, 9}

first:
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if j == 3 {
				break first
			}
			fmt.Println("i,", i, ",j,", j)
		}
	}
}

//goto
func Test_Goto_10(t *testing.T) {
	var a = 10
	//循环
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Println(a)
		a++
	}
}

//统计出一个文件里单词出现的频率。 分支循坏
//石头剪刀布
//{
//系统产生1-3的随机数，分别代表剪刀，石头和布
// 玩家键盘输入1-3数字，分别代表剪刀，石头和布
//       }
