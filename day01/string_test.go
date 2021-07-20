package day01

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func Test_String_01(t *testing.T) {
	var str = "我要好好的学习\n 冲冲冲"
	//fmt.Println(str)
	str = "我要好好的学习 \n \r 冲冲冲"
	fmt.Println(str)
}

func Test_String_02(t *testing.T) {
	var str1 = "xbb"
	var str2 = "xbb"
	fmt.Println(str1 >= str2)
}

func Test_String_03(t *testing.T) {
	var str1 = "我我我我我我我我"
	//每个字符串的下标获取 ASCII 字符
	for i := 0; i < len(str1); i++ {
		fmt.Printf("ascii: %c   %d \n  ", str1[i], str1[i])
	}
	//按Unicode字符遍历字符串
	for _, v := range str1 {
		fmt.Printf("unicode: %c   %d \n  ", v, v)
	}
}

func Test_String_04(t *testing.T) {
	var str1 = "我我我我我我"
	fmt.Println(len(str1))

	fmt.Println(utf8.RuneCountInString(str1))
	str1 = "aaaaaa"
	fmt.Println(len(str1))
}

func Test_String_05(t *testing.T) {
	var str = "hello,world!"
	//str1 := strings.Index(str, ",")
	//pos := strings.Index(str[str1:], "w")
	//fmt.Println(str1, pos, str[str1+pos:])
	str2 := strings.LastIndex(str, "w")
	fmt.Println(str2)
}

func Test_String_06(t *testing.T) {
	var str = "hello,world!"
	str1 := []byte(str)
	for i := 0; i < len(str1); i++ {
		str1[i] = ' '
	}
	fmt.Println(string(str1))
}

//拼接
func Test_String_07(t *testing.T) {
	var str = "hello,"
	var str1 = "world!"
	//声明缓冲区
	var stringBuilder bytes.Buffer
	//把字符串写入缓冲区
	stringBuilder.WriteString(str)
	stringBuilder.WriteString(str1)
	fmt.Println(stringBuilder.String())
}

//sprintf
func Test_String_08(t *testing.T) {
	var a int = 1
	var b int = 2
	res := fmt.Sprintf("我已经学了%d科了,还差%d科", a, b)
	fmt.Println(res)

	pi := 3.1415926
	res = fmt.Sprintf("%v %v %v", pi, true, "hello")
	fmt.Println(res)

	//声明匿名结构体
	stu := &struct {
		Name string
		Age  int
	}{
		Name: "zxz",
		Age:  20,
	}

	fmt.Println(fmt.Sprintf("使用'%%+v' %+v\n", stu))
	fmt.Println(fmt.Sprintf("使用'%%#+v' %#v\n", stu))
	fmt.Println(fmt.Sprintf("使用'%%T' %T\n", stu))
}

//base64
func Test_String_09(t *testing.T) {
	var str = "hello,world!"
	//把字符串编码
	encodeMsg := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encodeMsg)
	//解码
	res, err := base64.StdEncoding.DecodeString(encodeMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res))
	}
}

func Test_Common_String(t *testing.T) {
	/*
	   fmt,输入，输出
	       Println()
	       Print()
	       Printf()
	       Scanln()
	       Scanf()
	       ...
	   math,数学，
	       Abs()
	       Sqrt()
	       Pow(a,b)
	       Pow10(n)
	       ..
	   math/rand,
	       随机数
	       rand.Seed()
	       rand.Intn()
	   time，时间
	       time.Now()
	       Unix()
	       UnixNano()
	   strings包,字符串的常用的方法。。
	   包含：
	       Contains(s,substr)
	       ContainsAny(s,substr)
	   前缀和后缀
	       HasPrefix()
	       HasSuffix()
	   搜索：
	       Index(s,substr)
	       LastIndex(s,substr)
	   替换：
	       Replace(s,old,new,n),n表示替换的次数，如果全部替换，-1
	   统计：
	       Count(s,substr)
	   大小写转换：
	       ToLower()
	       ToUpper()
	   切割：
	       Split()
	       SplitN()
	   重复：
	       Repeat(s,n)
	   截取：
	       s[start:end]-->截取子串，[start,end)

	*/
	s1 := "HelloWorld"
	//1.是否包含指定的内容：
	fmt.Println(strings.Contains(s1, "hello"))    //判断s1字符串中是否包含指定的内容，返回值是bool类型
	fmt.Println(strings.ContainsAny(s1, "abcde")) //判断是否包含chars里的任意一个字符即可
	//2.搜索
	fmt.Println(strings.Index(s1, "l"))     //返回子串第一次出现的下标索引位置，如果子串不存在，返回-1。
	fmt.Println(strings.LastIndex(s1, "l")) //返回子串的最后一次出现的位置，
	//3.前缀，后缀
	s2 := "课堂笔记.txt"
	fmt.Println(strings.HasPrefix(s2, "2018")) //是否是以指定的内容开头
	fmt.Println(strings.HasSuffix(s2, ".txt")) //是否以指定的内容结束

	//4.统计
	fmt.Println(strings.Count(s1, "lloo")) //统计子串的次数
	//5.切割
	s3 := "rose,jack,jerry,,王二狗,,,"
	arr := strings.Split(s3, ",")
	fmt.Println(arr) //[[rose jack jerry 王二狗]]
	fmt.Println(len(arr))
	arr2 := strings.Split(s3, "")
	fmt.Println(arr2)

	arr3 := strings.SplitN(s3, ",", -1) //指定切割的个数，不超过n个。全切n=-1
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	//6.大小写转换
	s4 := "hello123WorLD"
	fmt.Println(strings.ToLower(s4))
	fmt.Println(strings.ToUpper(s4))

	//7.替换
	s5 := "hello world"
	fmt.Println(s5)
	fmt.Println(strings.Replace(s5, "l", "*", -1))

	//8.去除指定内容：首尾
	s6 := "+**zhang san   ***+"
	fmt.Println(s6)
	fmt.Println(strings.Trim(s6, "*+")) //去除首尾的指定内容
	fmt.Println(strings.TrimSpace(s6))  //去除首尾空格

	// 9.重复
	fmt.Println(strings.Repeat("hello", 5))

	//10.字符串的截取
	s7 := "helloworld" //substring(start,end)-->[start,end)
	s8 := s7[:5]
	fmt.Println(s8)
	fmt.Println(s7[2:7])
	fmt.Println(s7[5:])

}
