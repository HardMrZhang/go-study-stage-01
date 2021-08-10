package day07

import (
	"fmt"
	"testing"
)

type Cat struct {
}

type Dog struct {
}

type Speaker interface {
	speaker() string
}

func (d *Dog) speaker() string {
	return "狗"
}
func (c Cat) speaker() string {
	return "猫"
}
func Test_Interface_001(t *testing.T) {
	var s Speaker
	c := Cat{}
	d := &Dog{}
	s = c
	fmt.Println(s.speaker())
	s = d
	fmt.Println(s.speaker())
}

type Person struct {
}

type People interface {
	sleep()
	walk()
}

func (p Person) sleep() {
	fmt.Println("睡觉")
}
func (p Person) walk() {
	fmt.Println("走路")
}

func Test_Interface_002(t *testing.T) {
	var a interface{}
	a = 20
	v, ok := a.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("fail")
	}
	switch v := a.(type) {
	case int:
		fmt.Println("int", v)

	default:
		fmt.Println("sting", v)
	}

}

type Sleep interface {
}
type Walk interface {
}

type CJQ interface {
	Sleep
	Walk
}
