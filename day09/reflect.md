# go中的反射reflect

##### 什么是反射

在计算机科学中，反射是指计算机程序在运行时可以访问、检测和修改它本身状态或行为的一种能力。用比喻来说，反射就是程序在运行的时候能够“观察”并且修改自己的行为。

##### 什么情况需要使用反射

1. 使用场景
   - 不能明确接口调用哪个函数，需要根据传入的参数在运行时决定。

   - 不能明确传入函数的参数类型，需要在运行时处理任意对象。

2. 不推荐使用反射
   - 与反射相关的代码，经常是难以阅读的。在软件工程中，代码可读性也是一个非常重要的指标。
   - Go 语言作为一门静态语言，编码过程中，编译器能提前发现一些类型错误，但是对于反射代码是无能为力的。所以包含反射相关的代码，很可能会运行很久，才会出错，这时候经常是直接 panic，可能会造成严重的后果。
   - 反射对性能影响还是比较大的，比正常代码运行速度慢一到两个数量级。所以，对于一个项目中处于运行效率关键位置的代码，尽量避免使用反射特性。

##### go语言如何实现反射

`interface`，它是 Go 语言实现抽象的一个非常强大的工具。当向接口变量赋予一个实体类型的时候，接口会存储实体的类型信息，反射就是通过接口的类型信息实现的，反射建立在类型的基础上。

Go 语言在 reflect 包里定义了各种类型，实现了反射的各种函数，通过它们可以在运行时检测类型的信息、改变类型的值。

1. types和interface

   go中每个变量都有一个静态类型,在编译阶段就确定了的,这个类型是声明时候的类型,不是底层的数据结构

   ```go
   type MyInt int
   
   var i int
   var j MyInt //静态类型是MyInt
   ```

   尽管i,j底层类型都是int,但是他们是不同的静态类型,除非进行类型转换,不然i,j不能同时出现等号的左右两边

   **反射主要与interface{}类型相关。**

2. Interface{}

   ```go
   type iface struct {
   	tab  *itab
   	data unsafe.Pointer
   }
   
   type itab struct {
   	inter  *interfacetype //表示具体类型实现的接口类型
   	_type  *_type //表示具体类型
   	link   *itab
   	hash   uint32
   	_ [4]byte
   	fun    [1]uintptr
   }
   ```

   itab由具体类型`_type`以及`interfacetype`组成

   ![image-20210816162405592](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210816162405592.png)

   实际上,iface描述的是非空接口,他包含方法。e与之相对的eface描述的是空接口,不包含任何方法,

   **go语言里有的类型都实现了空接口**

   ```go
   type eface struct {
       _type *_type
       data  unsafe.Pointer
   }
   ```

   相比iface,eface比较简单了,只维护了一个_type字段,表示空接口所承载的具体的实体类型,data表示具体的值。

   ![image-20210816162959018](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210816162959018.png)

   ***接口变量可以存储任何实现了接口定义的所有方法的变量***

   ```go
   //go中最常见的就是 Reader 和 Writer 接口
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
   
   type Writer interface {
       Write(p []byte) (n int, err error)
   }
   ```

   ```go
   //接下来就是接口之间的转换和赋值
   var r io.Reader
   tty, err := os.OpenFile("/Users/zhangchaoyin/IdeaProjects/go-study-stage-01/day09/reflect.mds", os.O_RDWR, 0)
   if err != nil {
       return nil, err
   }
   r = tty
   ```

   首先,声明r的类型是io.Reader,注意,这是r的静态类型,此时动态类型为nil,动态值也为nil

   然后,将data赋值给r,此时r的动态类型为*os.File,动态值为非空,表示打开的文件对象<tty,*os.File>

   ![image-20210816164542776](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210816164542776.png)

   此时上图虽然fun指向一个read函数,但是 `*os.File`还包含writer函数,也就是说 `*os.File`还是实现了io.Write接口。

   ```go
   var w io.Writer
   w = r.(io.Writer)
   ```

   之所以用断言,而不能直接赋值,是因为 `r` 的静态类型是 `io.Reader`，并没有实现 `io.Writer` 接口。断言能否成功，看 `r` 的动态类型是否符合要求。

   w 也可以表示成 `<tty, *os.File>`，仅管它和 `r` 一样，但是 w 可调用的函数取决于它的静态类型 `io.Writer`，也就是说它只能有这样的调用形式: `w.Write()` 

   ![image-20210816165656269](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210816165656269.png)

   和 `r` 相比，仅仅是 `fun` 对应的函数变了：`Read -> Write`。

   最后我们测试一下空接口赋值进去是个什么效果

   ```go
   var empty interface{}
   empty = w
   //因为empty是空接口,因此所有类型都实现了他,所以直接赋值就行了
   ```

   ![image-20210816165940020](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210816165940020.png)

   从上面的三张图可以看到，interface 包含三部分信息：`_type` 是类型信息，`*data` 指向实际类型的实际值，`itab` 包含实际类型的信息，包括大小、包路径，还包含绑定在类型上的各种方法（图上没有画出方法），补充一下关于 os.File 结构体的图：

   ![image-20210816170311952](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210816170311952.png)

3. 反射的基本函数

   reflect包里定义了一个接口和结构体,reflect.Type和reflect.Value,他们提供了很多函数来获取存储在接口里的类型信息。

   - `reflect.Type` 主要提供关于类型相关的信息，所以它和 `_type` 关联比较紧密
   - `reflect.Value` 则结合 `_type` 和 `data` 两者，可以获取甚至改变类型的值。

   reflect 包中提供了两个基础的关于反射的函数来获取上述的接口和结构体:

   ```go
   func TypeOf(i interface{}) Type 
   func ValueOf(i interface{}) Value
   ```

   `TypeOf` 函数用来提取一个接口中值的类型信息。由于它的输入参数是一个空的 `interface{}`，调用此函数时，实参会先被转化为 `interface{}`类型。这样，实参的类型信息、方法集、值信息都存储到 `interface{}` 变量里了。

   ```go
   func TypeOf(i interface{}) Type {
   	eface := *(*emptyInterface)(unsafe.Pointer(&i))
   	return toType(eface.typ)
   }
   ```

   这里的 `emptyInterface` 和上面提到的 `eface` 是一回事（字段名略有差异，字段是相同的），并且在不同的源码包：前者在 `reflect` 包，后者在 `runtime` 包。 `eface.typ` 就是动态类型。

   ```go
   // emptyInterface is the header for an interface{} value.
   type emptyInterface struct {
   	typ  *rtype
   	word unsafe.Pointer
   }
   ```

   注意，返回值 `Type` 实际上是一个接口，定义了很多方法，用来获取类型相关的各种信息，而 `*rtype` 实现了 `Type` 接口。

   ```go
   type Type interface {
       // 所有的类型都可以调用下面这些函数
   
   	// 此类型的变量对齐后所占用的字节数
   	Align() int
   	
   	// 如果是 struct 的字段，对齐后占用的字节数
   	FieldAlign() int
   
   	// 返回类型方法集里的第 `i` (传入的参数)个方法
   	Method(int) Method
   
   	// 通过名称获取方法
   	MethodByName(string) (Method, bool)
   
   	// 获取类型方法集里导出的方法个数
   	NumMethod() int
   
   	// 类型名称
   	Name() string
   
   	// 返回类型所在的路径，如：encoding/base64
   	PkgPath() string
   
   	// 返回类型的大小，和 unsafe.Sizeof 功能类似
   	Size() uintptr
   
   	// 返回类型的字符串表示形式
   	String() string
   
   	// 返回类型的类型值
   	Kind() Kind
   
   	// 类型是否实现了接口 u
   	Implements(u Type) bool
   
   	// 是否可以赋值给 u
   	AssignableTo(u Type) bool
   
   	// 是否可以类型转换成 u
   	ConvertibleTo(u Type) bool
   
   	// 类型是否可以比较
   	Comparable() bool
   
   	// 下面这些函数只有特定类型可以调用
   	// 如：Key, Elem 两个方法就只能是 Map 类型才能调用
   	
   	// 类型所占据的位数
   	Bits() int
   
   	// 返回通道的方向，只能是 chan 类型调用
   	ChanDir() ChanDir
   
   	// 返回类型是否是可变参数，只能是 func 类型调用
   	// 比如 t 是类型 func(x int, y ... float64)
   	// 那么 t.IsVariadic() == true
   	IsVariadic() bool
   
   	// 返回内部子元素类型，只能由类型 Array, Chan, Map, Ptr, or Slice 调用
   	Elem() Type
   
   	// 返回结构体类型的第 i 个字段，只能是结构体类型调用
   	// 如果 i 超过了总字段数，就会 panic
   	Field(i int) StructField
   
   	// 返回嵌套的结构体的字段
   	FieldByIndex(index []int) StructField
   
   	// 通过字段名称获取字段
   	FieldByName(name string) (StructField, bool)
   
   	// FieldByNameFunc returns the struct field with a name
   	// 返回名称符合 func 函数的字段
   	FieldByNameFunc(match func(string) bool) (StructField, bool)
   
   	// 获取函数类型的第 i 个参数的类型
   	In(i int) Type
   
   	// 返回 map 的 key 类型，只能由类型 map 调用
   	Key() Type
   
   	// 返回 Array 的长度，只能由类型 Array 调用
   	Len() int
   
   	// 返回类型字段的数量，只能由类型 Struct 调用
   	NumField() int
   
   	// 返回函数类型的输入参数个数
   	NumIn() int
   
   	// 返回函数类型的返回值个数
   	NumOut() int
   
   	// 返回函数类型的第 i 个值的类型
   	Out(i int) Type
   
       // 返回类型结构体的相同部分
   	common() *rtype
   	
   	// 返回类型结构体的不同部分
   	uncommon() *uncommonType
   }
   ```

   注意到 `Type` 方法集的倒数第二个方法 `common` 返回的 `rtype`类型，它和上面讲到的 `_type` 是一回事，而且源代码里也注释了:两边要保持同步

   ```
   type rtype struct {
   	size       uintptr
   	ptrdata    uintptr
   	hash       uint32
   	tflag      tflag
   	align      uint8
   	fieldAlign uint8
   	kind       uint8
   	alg        *typeAlg
   	gcdata     *byte
   	str        nameOff
   	ptrToThis  typeOff
   }
   ```

   所有的类型都会包含 `rtype` 这个字段，表示各种类型的公共信息；另外，不同类型包含自己的一些独特的部分。

   `Type` 接口实现了 `String()` 函数，满足 `fmt.Stringer` 接口，因此使用 `fmt.Println` 打印的时候，输出的是 `String()` 的结果。另外，`fmt.Printf()` 函数，如果使用 `%T` 来作为格式参数，输出的是 `reflect.TypeOf` 的结果，也就是动态类型。

   讲完了 `TypeOf` 函数，再来看一下 `ValueOf` 函数。返回值 `reflect.Value` 表示 `interface{}` 里存储的实际变量，它能提供实际变量的各种信息。相关的方法常常是需要结合类型信息和值信息。例如，如果要提取一个结构体的字段信息，那就需要用到 _type (具体到这里是指 structType) 类型持有的关于结构体的字段信息、偏移信息，以及 `*data` 所指向的内容 —— 结构体的实际值。

   ```go
   func ValueOf(i interface{}) Value {
   	if i == nil {
   		return Value{}
   	}
   
   	// TODO: Maybe allow contents of a Value to live on the stack.
   	// For now we make the contents always escape to the heap. It
   	// makes life easier in a few places (see chanrecv/mapassign
   	// comment below).
   	escapes(i)
   
   	return unpackEface(i)
   }
   //分解eface
   func unpackEface(i interface{}) Value {
   	e := (*emptyInterface)(unsafe.Pointer(&i))
   	// NOTE: don't read e.word until we know whether it is really a pointer or not.
   	t := e.typ
   	if t == nil {
   		return Value{}
   	}
   	f := flag(t.Kind())
   	if ifaceIndir(t) {
   		f |= flagIndir
   	}
   	return Value{t, e.word, f}
   }
   ```

   源码解析:将先将 `i` 转换成 `*emptyInterface` 类型， 再将它的 `typ` 字段和 `word` 字段以及一个标志位字段组装成一个 `Value` 结构体，而这就是 `ValueOf` 函数的返回值，它包含类型结构体指针、真实数据的地址、标志位。

   Value 结构体定义了很多方法，通过这些方法可以直接操作 Value 字段 ptr 所指向的实际数据:

   ```go
   // 设置切片的 len 字段，如果类型不是切片，就会panic
    func (v Value) SetLen(n int)
    
    // 设置切片的 cap 字段
    func (v Value) SetCap(n int)
    
    // 设置字典的 kv
    func (v Value) SetMapIndex(key, val Value)
   
    // 返回切片、字符串、数组的索引 i 处的值
    func (v Value) Index(i int) Value
    
    // 根据名称获取结构体的内部字段值
    func (v Value) FieldByName(name string) Value
   
    // 用来获取 int 类型的值
    func (v Value) Int() int64
   
    // 用来获取结构体字段（成员）数量
    func (v Value) NumField() int
   
    // 尝试向通道发送数据（不会阻塞）
    func (v Value) TrySend(x reflect.Value) bool
   
    // 通过参数列表 in 调用 v 值所代表的函数（或方法
    func (v Value) Call(in []Value) (r []Value) 
   
    // 调用变参长度可变的函数
    func (v Value) CallSlice(in []Value) []Value 
   ```

   另外，通过 `Type()` 方法和 `Interface()` 方法可以打通 `interface`、`Type`、`Value` 三者。Type() 方法也可以返回变量的类型信息，与 reflect.TypeOf() 函数等价。Interface() 方法可以将 Value 还原成原来的 interface。

   ![image-20210816175542415](/Users/zhangchaoyin/Library/Application Support/typora-user-images/image-20210816175542415.png)

4. 总结

   - `TypeOf()` 函数返回一个接口，这个接口定义了一系列方法，利用这些方法可以获取关于类型的所有信息；
   - `ValueOf()` 函数返回一个结构体变量，包含类型信息以及实际值。

##### 反射三大定律

1. 反射是一种检测存储在 `interface` 中的类型和值机制。这可以通过 `TypeOf` 函数和 `ValueOf` 函数得到。

2. 将 `ValueOf` 的返回值通过 `Interface()` 函数反向转变成 `interface` 变量。

3. 如果需要操作一个反射变量，那么它必须是可设置的

   ```
   前两条就是说 接口型变量 和 反射类型对象 可以相互转化，反射类型对象实际上就是指的前面说的 reflect.Type 和 reflect.Value。
   
   第三条不太好懂,反射变量可设置的本质是它存储了原变量本身，这样对反射变量的操作，就会反映到原变量本身；反之，如果反射变量不能代表原变量，那么操作了反射变量，不会对原变量产生任何影响，这会给使用者带来疑惑。所以第二种情况在语言层面是不被允许的。
   
   如果想要操作原变量，反射变量 Value 必须要对原变量的地址有把握才行。
   ```

   

##### 反射有哪些应用

- IDE 中的代码自动补全功能
- 对象序列化（encoding/json）
- fmt 相关函数的实现
- ORM（全称是：Object Relational Mapping，对象关系映射）

##### 如何比较两个对象完全相同

