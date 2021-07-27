package main

import (
	"fmt"
	"testing"
)

func f1() {
	fmt.Println("f1")
}

type Test struct {
	Name string
}

func (t *Test) Close() {
	fmt.Println(t.Name, "close")
	//return t.Name
}
func Test_func_0001(t *testing.T) {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, v := range ts {
		// a , b,c
		//res :=v.Close()
		//fmt.Println(v.Close())
		v1 := v
		defer v1.Close()
	}

}
