package day03

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_list_01(t *testing.T) {
	l := list.New()
	l.PushBack("zzy")
	element := l.PushFront(20)
	// *list.Element 结构，这个结构记录着列表元素的值以及与其他节点之间的关系等信息
	fmt.Println(element)
	// 在fist之后添加high
	l.InsertAfter("cz", element)
	l.Remove(element)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func Test_list_02(t *testing.T) {
	l := list.New()
	l.PushBack("zzy")
	l.PushFront(20)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
