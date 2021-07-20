package day01

import (
	"fmt"
	"testing"
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
