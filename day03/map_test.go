package day03

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

func Test_Map_01(t *testing.T) {
	var m1 map[string]int
	var m2 map[string]int
	m1 = map[string]int{"one": 1, "two": 2}
	m3 := make(map[string]float32, 100)
	//map赋值
	m2 = m1

	m3["grade"] = 3.1

	//m2 是 m1 的引用，对 m2 的修改也会影响到 m1 的值
	m2["one"] = 3
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3["grade"])

	//不能使用new声明
	//m4 := new(map[string]float64)
	//m4["one"] = 1.11
	//fmt.Println("m4:", m4)

	//可以用切片作为map的值
	mp1 := make(map[string][]int)
	mp2 := make(map[string]*[]int)
	arr := []int{1, 2, 3, 4, 5}
	mp1["one"] = arr[1:3]
	tmp := arr[1:3]
	mp2["two"] = &tmp

	//value可以是map的情况
	stu := make(map[string]map[string]string)

	stu["stu01"] = make(map[string]string)
	stu["stu01"]["name"] = "zhangsan"
	stu["stu01"]["sex"] = "男"
	fmt.Println(stu)
}

func Test_Map_02(t *testing.T) {
	var m1 = make(map[string]int)
	m1["c"] = 1
	m1["a"] = 2
	m1["b"] = 3
	//如果只查询k,v可以直接省略
	for k, v := range m1 {
		fmt.Printf("key : %v , value : %v \n", k, v)
	}
	//如果期望map按照特定的顺序返回结果
	var s []string
	for k := range m1 {
		s = append(s, k)
	}
	sort.Strings(s)
	fmt.Println(s)
}

func Test_Map_03(t *testing.T) {
	var m1 = make(map[string]int)
	m1["c"] = 1
	m1["a"] = 2
	m1["b"] = 3
	//m1["a"] = 4
	delete(m1, "a")

	for k, v := range m1 {
		fmt.Printf("key : %v , value : %v \n", k, v)
	}

}

//map的安全性讨论
func Test_Map_04(t *testing.T) {
	var m = make(map[int]int)
	//开启并发处理
	go func() {
		for {
			//不断对m[0]进行赋值
			m[0] = 1
		}
	}()

	//再开启另一个并发处理
	go func() {
		for {
			//不断对m[0]进行读取
			_ = m[0]
		}
	}()

	/**
	并发的 map 读和 map 写，也就是说使用了两个并发函数不断地对 map 进行读和写而发生了竞态问题，
	map 内部会对这种并发操作进行检查并提前发现。
	*/

}
func Test_Map_05(t *testing.T) {
	var sync_map sync.Map

	//将键值对保存到sync_map中
	sync_map.Store("a", 1)
	sync_map.Store("b", 2)
	sync_map.Store("c", 3)
	//读取
	fmt.Println(sync_map.Load("a"))
	//删除a
	sync_map.Delete("a")
	//遍历
	sync_map.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	fmt.Println(sync_map.Load("a"))
}

func Test_Map_06(t *testing.T) {

	//var m1 map[string]int
	var m2 sync.Map
	m2.Store("a", 1)
	m2.Delete("a")
	fmt.Println(m2.Load("a"))
	m2.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
