# go语言学习大纲

------

### Go语言基础语法

1. go语言特点和优势

   - 快速编译，高效执行，易于开发。

   - 对于网络通信、并发和并行编程有很好的支持，可以更好地利用大量的分布式和多核的计算机。

   - Go语言通过 goroutine 这种轻量级线程的概念来实现这个目标，然后通过 channel 来实现各个 goroutine 之间的通信。它们实现了分段栈增长和 goroutine 在线程基础上多路复用技术的自动化。

   - Go语言从本质上（程序和结构方面）来实现并发编程。

   - Go语言作为强类型语言，隐式的类型转换是不被允许的。一条原则：让所有的东西都是显式的。

   - Go语言本身是由C语言开发的，而不是Go语言（Go1.5开始自举）。

   - Go语言的二进制文件体积是最大的（每个可执行文件都包含 runtime）。

   - Go语言做为一门静态语言，但却和很多动态脚本语言一样得灵活。

     ```ABAP
     运行时(Runtime)是指将数据类型的确定由编译时推迟到了运行时
     Runtime是一套比较底层的纯C语言API, 属于1个C语言库, 包含了很多底层的C语言API
     ```

2. go适合做什么

   1. 服务端开发
   2. 分布式系统，微服务
   3. 网络编程 net/http
   4. 区块链开发
   5. 内存KV数据库，例如boltDB、levelDB
   6. 云平台,中台

3. 基础语法

   - [x] 例子

     ```golang
     package main
     
     import (
     	"fmt"
     )
     //init()
     func main() {
     	fmt.Println("Hello World!")
     }
     ```

     - 包声明：package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
       - 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
       - 引入包：使用import圆括号进行引入包。
     - 引入标准包
     - 引入第三方包

   - [x] 关键字:关键字是一些特殊的用来帮助编译器理解和解析源代码的单词。Go语言中有25个关键字或保留字。只能用在语法允许的地方，不能作为名称使用。

     1. 声明变量：const,func,var,import,package,type
     2. 组合类型:chan,interface,map,struct
     3. 流程控制语句:switch,case,return,break,default,fallthrough,for,range,continue,else,goto,if,select
     4. 特殊关键字:defer,go。也可以看作是流程控制关键字， 但它们有一些特殊的作用。

   - [ ] 预定义标识符：三十几个内置的预申明的常量、类型和函数。所有的类型名、变量名、常量名、跳转标签、包名和包的引入名都必须是标识符。

     1. 常量:true、flase、iota、nil

     2. 类型:int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、uintptr、float32、float64、complex128(实部和虚部都是float64)、complex64(实部和虚部都是float32)、bool、byte(uint8 的别名)、rune(int32 的别名 代表一个 Unicode 码)、string、error

     3. 内置函数函数:make、len、cap、new、append、copy、close、delete、complex、real、imag、panic、recover

        ```ABAP
        make: make函数是Go的内置函数，它的作用是为slice、map或chan初始化并返回引用。make仅仅用于创建slice、map和channel，并返回一个初始化的（不是零）的，type T的，不是 *T 的值。
        
        len: 来求长度，比如string、array、slice、map、channel ，返回长度
        
        cap: 返回的是数组切片分配的空间大小（只能用于切片和 map）
        
        new: new的作用是初始化一个指向类型的指针（*T）。使用new函数来分配空间，传递给new函数的是一个类型，不是一个值。返回的是指向这个新分配的零值的指针。
        
        append: 用来追加元素到数组、slice中,返回修改后的数组、slice
        
        copy: 用于复制和连接slice，返回复制的数目
        
        delete:从map中删除key对应的value
        
        complex:返回complex的虚部
        
        real:返回complex的实部complex、real imag：用于创建和操作复数）
        
        imag:返回complex的虚部
        
        panic: 停止常规的goroutine(panic和recover：用来做错误处理）
        
        recover:允许程序定义goroutine的panic动作
        
        close:主要用来关闭channel
        
        
        ```

   - [ ] 内置接口

     ```golang
     //只要实现了Error()函数，返回值为String的都实现了err接口
     type error interface { 
     Error() String
     }
     ```

   - [ ] 什么是变量

     - 变量是为存储特定类型的值而提供给内存的位置。

   - [ ] 声明变量

     - var名称类型是声明单个变量的语法

     ```golang
     //以字母或下划线开头，由一个或多个字母、数字、下划线组成
     var name type
     name = value
     ```

     - 根据值自行判定变量类型,如果一个变量有一个初始值，Go将自动能够使用初始值来推断该变量的类型。因此，如果变量具有初始值，则可以省略变量声明中的类型。

     ```
     var name = value
     ```

     - 省略var,注意 :=左侧的变量不应该是已经声明过的(多个变量同时声明时，至少保证一个是新变量)，否则会导致编译错误(简短声明)

     ```golang
     name:=value
     //例如
     //这种方式它只能被用在函数体内，而不可以用于全局变量的声明与赋值
     var a int = 10
     var b = 10
     c:=10
     ```

     - 多重变量声明

     ```golang
     //1.以逗号分隔，声明与赋值分开，若不赋值，存在默认值
     var n1,n2,n3 type
     n1,n2,n3 = v1,v2,v3
     //2.直接赋值,变量类型可以是不同的类型
     var name1, name2, name3 = v1, v2, v3
     //3.集合类型
     var (
     name1 type1
     name2 type2
     )
     var (
     a int
     b string
     c []float32
     d func() bool
     e struct {
     x int
     }
     )
     ```

     ```golang
     如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20 就是不被允许的，编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值。如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。func main() {   var a string = "abc"   fmt.Println("hello, world")}尝试编译这段代码将得到错误 a declared and not used----------------------------------------------------------------------------在同一个作用域中，已存在同名的变量，则之后的声明初始化，则退化为赋值操作。但这个前提是，最少要有一个新的变量被定义，且在同一作用域，例如，下面的y就是新定义的变量package mainimport ("fmt")func main() {x := 140fmt.Println(&x)x, y := 200, "abc"fmt.Println(&x, x)fmt.Print(y)}
     ```

   - [ ] 空白标识符 _ 也被用于抛弃值(匿名参数)，如值 5 在：_, b = 5, 7 中被抛弃

     - _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值

       并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：val, err = Func1(var1)

       1. 变量必须先声明，才能够使用，而且每个变量只能被声明一次。
       2. 因为go是强类型语言，赋值类型要对应 
       3. name := value，这种声明变量的方式，不能在函数外部使用
       4. 默认值：也叫零值。

   - [ ] 常量声明

     - 常量是一个简单值的标识符，在程序运行时，不会被修改的量。

     - 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型

     - 不曾使用的常量，在编译的时候，是不会报错的

     - 显示指定类型的时候，必须确保常量左右值类型一致，需要时可做显示类型转换。这与变量就不一样了，变量是可以是不同的类型值

     - 在函数体外声明的变量称之为全局变量，全局变量只需要在一个源文件中定义，就可以在所有源文件中使用，当然，不包含这个全局变量的源文件需要使用“import”关键字引入全局变量所在的源文件之后才能使用这个全局变量。

       ````
       const identifier [type] = value
       ````

       ```golang
       显式类型定义： const b string = "abc"隐式类型定义： const b = "abc"//常量组,可做枚举用const (	 Unknown = 0	 Female = 1	 Male = 2)
       ```

     - 特殊常量iota

       1. iota 在 const关键字出现时将被重置为 0，const 中每新增一行常量声明将使 iota 计数一次。
       2. iota可以理解为const语句块的行索引
       3. iota也可以用作枚举

       ```golang
       const (	a = iota //0	b = iota //1	c = iota //2)const (	a = iota //0b        //1c = "aa" //aad 			 //aae = iota //4_  f        //6) const ( a, b = iota + 1, iota + 2 //1,2 c, d                      //2,3 e, f                      //3,4 )//使用位左移与 iota 计数配合可优雅地实现存储单位的常量枚举func TestIota(t *testing.T) {	type ByteSize float64	const (		_           = iota // 通过赋值给空白标识符来忽略值		KB ByteSize = 1 << (10 * iota)		MB ByteSize = 1 << (10 * iota)		GB ByteSize = 1 << (10 * iota)		TB ByteSize = 1 << (10 * iota)	)	fmt.Println(KB)	fmt.Println(MB)	fmt.Println(GB)	fmt.Println(TB)}
       ```

     - [ ] 全局变量,局部变量,形式参数

     - [ ] Go中字符串操作

       - Go中的字符串是一个字节的切片。是不可概念的字节序列。可以通过将其内容封装在""中来创建字符串。Go中的字符串是Unicode兼容的，并且是UTF-8编码的。
       - 字符串是一种值类型(存贮在栈上)，且值不可变，即创建某个文本后将无法再次修改这个文本的内容，更深入地讲，字符串是字节的定长数组
       - 转义符

     - `\n`：换行符

     - `\r`：回车符

     - `\t`：tab 键

     - `\u 或 \U`：Unicode 字符

     - `\\`：反斜杠自身

       ```golang
       func Test_String_01(t *testing.T) {	var str = "我要好好的学习\n 冲冲冲"	//fmt.Println(str)	str = "我要好好的学习 \n \r 冲冲冲"	fmt.Println(str)}
       ```

       - 字符串的比较：一般的比较运算符`==、!=、<、<=、>=、>`是通过在内存中按字节比较来实现字符串比较的，因此比较的结果是字符串自然编码的顺序。

     1. 字符串所占的字节长度可以通过函数 `len()` 来获取，例如 `len(str)`。
        字符串的内容（纯字节）可以通过标准索引法来获取，在方括号[ ]内写入索引，索引从 0 开始计数,str[0] str[len(len)-1]

       - 字符串拼接

       ```
     //方案1str := "a"+"b"str1 := "a" + "b" str1 += "c"------------------------------------------------------------------------//方案2var str = "hello,"var str1 = "world!"//声明缓冲区var stringBuilder bytes.Buffer//把字符串写入缓冲区stringBuilder.WriteString(str)stringBuilder.WriteString(str1)fmt.Println(stringBuilder.String())
       ```

       - 定义多行字符串

       ```
     //使用双引号书写字符串的方式是字符串常见表达方式之一，被称为字符串字面量。这种双引号字面量不能跨行，如果想要在源码中嵌入一个多行字符串时，就必须使用\`反引号var str = `abc\r\n`
       ```

       - 字符串的遍历

     ```
     ASCII 字符串遍历直接使用下标。Unicode 字符串遍历用 `for range`。
     ```

       - 计算字符串的长度

     ```golang
     ASCII 字符串长度使用 len() 函数。Unicode 字符串长度使用 utf8.RuneCountInString() 函数。
     ```

       - 字符串的截取

     ```
     strings.Index：正向搜索子字符串。strings.LastIndex：反向搜索子字符串。搜索的起始位置可以通过切片偏移制作。
     ```

       - 修改字符串

     ```
     字符串默认是不可变的,字符串不可变有很多好处，如天生线程安全,大家使用的都是只读对象，无须加锁；再者，方便内存共享，而不必使用写时复制等技术；字符串 hash 值也只需要制作一份。所以说，代码中实际修改的是 []byte，[]byte 在 Go 语言中是可变的，本身就是一个切片。在完成了对[]byte操作后,再使用 string()将[]byte转为字符串时，重新创造了一个新的字符串。总结:- 修改字符串时，可以将字符串转换为[]byte进行修改。- []byte和string可以通过强制类型转换互转
     ```

       - fmt.Sprintf格式化输出

     ```
     | %v | 按原本值来输出                            | %+v| 在 `%v`的基础上对结构体字段名和值进行展开 | %#v| 输出 GO 语言语法格式的值                 | %T | 输出 GO 语言语法格式的类型                | %% | 输出 `%`本体                              | %b | 以二进制的方式显示                        | %o | 以八进制的方式显示                        | %x | 一十六进制的方式显示                      | %U | Unicode 字符                              | %f | 浮点型字符                                | %p | 指针,十六进制的方式显示                   
     ```

       - base64编码

     ```
     Base64编码是常见的对8比特字节码的编码方式之一。Base64可以使用64个可打印字符来表示二进制数据，电子邮件就是使用这种编码。
     ```

     - [x] 分支循环

       - 条件语句

     - if语句

     - 语法格式

       ```golang
       if 布尔表达式 {   /* 在布尔表达式为 true 时执行 */}if 布尔表达式 {   /* 在布尔表达式为 true 时执行 */} else {  /* 在布尔表达式为 false 时执行 */}if 布尔表达式1 {   /* 在布尔表达式1为 true 时执行 */} else if 布尔表达式2{   /* 在布尔表达式1为 false ,布尔表达式2为true时执行 */} else{   /* 在上面两个布尔表达式都为false时，执行*/}
       ```

       - 其中包含一个可选的语句组件(在评估条件之前执行)，则还有一个变体。

     ```
     if statement; condition {  	//Todo}if condition{ //Todo}
     ```

     - switch

       - switch是一个条件语句，它计算表达式并将其与可能匹配的列表进行比较，并根据匹配执行代码块。它可以被认为是一种惯用的方式来写多个if else子句。 

     - select

       - 循环语句

     - for

     - 跳出循环语句

     - goto

     - [x] 数组和切片

     - [ ] Map

     - [ ] 函数基础

     - [ ] 函数式编程

     - [ ] 指针

     - [ ] 结构体

     - [ ] 方法

     - [ ] 接口

     - [ ] 错误处理

     - [ ] 通道Channel

     - [ ] 网络编程

     - [ ] 反射Reflect

     - [ ] Common func

       

     

     ​	

     

     

