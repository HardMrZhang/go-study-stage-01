package day06

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

/**
二维矢量模拟玩家移动
在游戏中，一般使用二维矢量保存玩家的位置，使用矢量运算可以计算出玩家移动的位置，
实现二维矢量对象，接着构造玩家对象，最后使用矢量对象和玩家对象共同模拟玩家移动的过程。
*/
//矢量是数学中的概念，二维矢量拥有两个方向的信息，同时可以进行加、减、乘（缩放）、距离、单位化等计算，
//在计算机中，使用拥有 X 和 Y 两个分量的 Vec2 结构体实现数学中二维向量的概念，
type Vec2 struct {
	X, Y float32
}

//加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

//减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

//乘
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

//距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	//勾股定理   math.Sqrt:开根号
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

//定义矢量单位化
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec2{0, 0}
}

type Player struct {
	currPos   Vec2    // 当前位置
	targetPos Vec2    // 目标位置
	speed     float32 // 移动速度
}

// 移动到某个点就是设置目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

// 获取当前的位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

// 是否到达
func (p *Player) IsArrived() bool {
	// 通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经到达目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

// 逻辑更新
func (p *Player) Update() {
	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		// 移动完成后，更新当前位置
		p.currPos = newPos
	}
}

// 创建新玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}
func Test_struct_practise_01(t *testing.T) {
	// 实例化玩家对象，并设速度为0.5
	p := NewPlayer(0.5)
	// 让玩家移动到3,1点
	p.MoveTo(Vec2{3, 1})
	// 如果没有到达就一直循环
	for !p.IsArrived() {
		// 更新玩家位置
		p.Update()
		// 打印每次移动后的玩家位置
		fmt.Println(p.Pos())
	}
}

//使用事件系统实现事件的响应和处理
/**
一个事件系统拥有如下特性：
能够实现事件的一方，可以根据事件 ID 或名字注册对应的事件。
事件发起者，会根据注册信息通知这些注册者。
一个事件可以有多个实现方响应。
*/

//实例化一个通过字符串映射函数函数切片的map
var eventByName = make(map[string][]func(interface{}))

//注册事件,提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {
	//通过名字查找事件列表
	list := eventByName[name]
	//在列表切片添加函数
	list = append(list, callback)

	//将修改的时间列表切片保存回去
	eventByName[name] = list
}

//调用事件
func CallEvent(name string, param interface{}) {
	//通过名字找到时间列表
	list := eventByName[name]
	//遍历这个时间的所有回调
	for _, callback := range list {
		//传入参数调用回调
		callback(param)
	}
}

//声明角色的结构体
type Actor struct {
}

//为角色添加一个事件处理函数
func (a *Actor) onEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

//全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}
func Test_struct_practise_02(t *testing.T) {
	//实例化一个角色
	a := new(Actor)
	//注册 OnSkill的回调
	RegisterEvent("OnSkill", a.onEvent)

	//再次在OnSkill上注册全局事件
	RegisterEvent("OnSkill", GlobalEvent)
	//调用事件,所有注册的同名函数都会被调用
	CallEvent("OnSkill", 100)
	fmt.Println(eventByName)
}

/**
结构体内嵌模拟类的继承
*/

type Flying struct {
}

func (f *Flying) Fly() {
	fmt.Println("can fly")
}

type Walkable struct {
}

func (w *Walkable) Walk() {
	fmt.Println("can walk")
}

type Human struct {
	Walkable
}
type Bird struct {
	Walkable
	Flying
}

func Test_struct_practise_03(t *testing.T) {
	b := new(Bird)
	b.Fly()
	b.Walk()

	h := new(Human)
	h.Walk()
}

/**
使用匿名结构体解析JSON数据
JavaScript 对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。

Go语言对于这些标准格式的编码和解码都有良好的支持，由标准库中的 encoding/json、encoding/xml、encoding/asn1 等包提供支持，并且这类包都有着相似的 API 接口。

基本的 JSON 类型有数字（十进制或科学记数法）、布尔值（true 或 false）、字符串，其中字符串是以双引号包含的 Unicode 字符序列，支持和Go语言类似的反斜杠转义特性，不过 JSON 使用的是 \Uhhhh 转义数字来表示一个 UTF-16 编码，而不是Go语言的 rune 类型。

*/

type A struct {
	Ax     float32
	Ay, Az int
}

type B struct {
	Bx int
}

func genJsonData() []byte {
	raw := &struct {
		A
		B
		C bool
	}{

		A: A{
			Ax: 5.5,
			Ay: 1920,
			Az: 1080,
		},

		B: B{
			2910,
		},

		C: true,
	}
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func Test_struct_practise_04(t *testing.T) {
	jsonData := genJsonData()
	fmt.Println(string(jsonData))
	// 只需要屏幕和指纹识别信息的结构和实例
	AAndC := struct {
		A
		C bool
	}{}
	// 反序列化到screenAndTouch
	json.Unmarshal(jsonData, &AAndC)
	// 输出AAndC的详细结构
	fmt.Printf("%+v\n", AAndC)
	// 只需要A和C信息的结构和实例
	BAndC := struct {
		B
		C bool
	}{}
	// 反序列化到BAndC
	json.Unmarshal(jsonData, &BAndC)
	// 输出BAndC的详细结构
	fmt.Printf("%+v\n", BAndC)
}

/**
结构体数据保存为JSON格式数据
*/
func Test_struct_practise_05(t *testing.T) {
	type Skill struct {
		Name  string `json:"name,omitempty"` // omitempty，来过滤掉转换的JSON格式中的空值
		Level int    `json:"level"`
	}
	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}
	a := Actor{
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "", Level: 1},
			{Name: "hello"},
			{Name: "world", Level: 3},
		},
	}

	res, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	jsonData := string(res)
	fmt.Println(jsonData)
}
