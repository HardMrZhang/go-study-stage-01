package day06

import (
	"fmt"
	"testing"
)

/**

在Go语言中，结构体就像是类的一种简化形式，那么类的方法在哪里呢？在Go语言中有一个概念，它和方法有着同样的名字，并且大体上意思相同，Go 方法是作用在接收器（receiver）上的一个函数，接收器是某种类型的变量，因此方法是一种特殊类型的函数。

接收器类型可以是（几乎）任何类型，不仅仅是结构体类型，任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型，但是接收器不能是一个接口类型，因为接口是一个抽象定义，而方法却是具体实现，如果这样做了就会引发一个编译错误invalid receiver type…。

接收器也不能是一个指针类型，但是它可以是任何其他允许类型的指针，一个类型加上它的方法等价于面向对象中的一个类，一个重要的区别是，在Go语言中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在不同的源文件中，唯一的要求是它们必须是同一个包的。

类型 T（或 T）上的所有方法的集合叫做类型 T（或 T）的方法集。

因为方法是函数，所以同样的，不允许方法重载，即对于一个类型只能有一个给定名称的方法，但是如果基于接收器类型，是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收器类型上存在，比如在同一个包里这么做是允许的。
提示：
	在面向对象的语言中，类拥有的方法一般被理解为类可以做的事情。在Go语言中“方法”的概念与其他语言一致，只是Go语言建立的“接收器”强调方法的作用对象是接收器，也就是类实例，而函数没有作用对象。
*/

//为结构体添加方法
type Bag struct {
	item []int
}

//模拟将物品放入背包的过程
func Insert(b *Bag, itemId int) {
	b.item = append(b.item, itemId)
}
func Test_struct_sec_01(t *testing.T) {
	bag := new(Bag)
	Insert(bag, 20)
}

//(b*Bag) 表示接收器，即 Insert 作用的对象实例
func (b *Bag) Insert02(itemId int) {
	b.item = append(b.item, itemId)
}
func Test_struct_sec_02(t *testing.T) {
	bag := new(Bag)
	Insert(bag, 20)
}

//指针类型的接收器
/**
指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的 this 或者 self。

由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的。
*/
type Property struct {
	value int
}

func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) GetValue() int {
	return p.value
}

func Test_struct_sec_03(t *testing.T) {
	p := new(Property)
	p.SetValue(20)
	fmt.Println(p.GetValue())
}

//非指针类型的接收器
/**
当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效。
*/

type Coordinate struct {
	X int
	Y int
}

func (c Coordinate) add(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X + other.X,
		Y: c.X + other.Y,
	}
}
func Test_struct_sec_04(t *testing.T) {
	c1 := Coordinate{
		X: 1,
		Y: 2,
	}
	c2 := Coordinate{
		X: 3,
		Y: 4,
	}
	r := c1.add(c2)
	fmt.Println(r)
}

/**
Go语言可以将类型的方法与普通函数视为一个概念，从而简化方法和函数混合作为回调类型时的复杂性。

调用者无须关心谁来支持调用，系统会自动处理是否调用普通函数或类型的方法。

无论是普通函数还是结构体的方法，只要它们的签名一致，与它们签名一致的函数变量就可以保存普通函数或是结构体方法。
*/

type Person struct {
	eye string
}

//person类
func (p *Person) sleep(w string) {}

func Test_struct_sec_0100(t *testing.T) {
	p := new(Person)
	p.sleep("")

}

type class struct{}

func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}
func funcDo(v int) {
	fmt.Println("call function do ", v)
}

func Test_struct_sec_05(t *testing.T) {
	//声明一个函数回调
	var delegate func(int)
	c := new(class)
	//将回调设为c的do方法
	delegate = c.Do
	delegate(10)

	//将回调设为普通函数
	delegate = funcDo

	delegate(100)

}
