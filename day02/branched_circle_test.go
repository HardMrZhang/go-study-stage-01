package day02

//go tool compile -N -l -S

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
	"unicode"
	"unicode/utf8"
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

/**
case中的表达式是可选的，可以省略。如果该表达式被省略，则被认为是switch true，
并且每个case表达式都被计算为true，并执行相应的代码块。
*/
func Test_Switch_04(t *testing.T) {
	num := 10
	switch {
	case num >= 0 && num <= 50:
		fmt.Println("d")
	case num > 50 && num <= 100:
		fmt.Println("c")
	}
}

func Test_Switch_05(t *testing.T) {
	/*
	   switch的其他写法：
	   1.省略switch后的表达式，相当于直接作用在true
	   2.case后可以同时匹配多个数据,匹配上任意一个都可以执行该case分支
	   3.switch后支持多一条初始化语句
	*/

	//1.省略switch后的表达式
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}

	//2.case后有多个数值
	var str = "a"
	switch str {
	case "a", "b", "c", "d":
		fmt.Println("在其中")
	default:
		fmt.Println("不在其中")
	}

	//3.多一条初始化语句
	switch str1 := "golang"; str1 {
	case "java":
		fmt.Println("java")
	case "C++":
		fmt.Println("c++")
	default:
		fmt.Println("R")

	}
}

//判断type
func Test_Switch_06(t *testing.T) {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x的类型 ：%T", i)
	case int:
		fmt.Printf("x的类型是int")
	case float64:
		fmt.Printf("x的类型是float64")
	case func(int):
		fmt.Printf("x是func(int)型")
	case bool, string:
		fmt.Printf("x是bool或string型")
	default:
		fmt.Println("未知类型")

	}
}

//select
func Test_Switch_07(t *testing.T) {
	ch := make(chan int)
	quit := make(chan bool)
	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}() //别忘了()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	<-quit
	fmt.Println("程序结束")
}

func Test_For_08(t *testing.T) {
	var a = 5
	var b = 6
	nums := [10]int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("数组元素为 %d\n", nums[i])
	}
	for a > b {
		fmt.Println("a>b ")
	}
	//break
	for k, _ := range nums {
		if k > 3 {
			fmt.Println(nums[k])
			break
		}
	}
	//continue
	for _, v := range nums {
		if v%2 == 0 {
			fmt.Printf("能被二整除的是%d \n ", v)
			continue
		}
	}

}

//贴标签
func Test_For_09(t *testing.T) {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	var b = [5]int{1, 3, 5, 7, 9}

first:
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if j == 3 {
				break first
			}
			fmt.Println("i,", i, ",j,", j)
		}
	}
}

//goto
func Test_Goto_10(t *testing.T) {
	var a = 10
	//循环
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Println(a)
		a++
	}
}

//统计出一个文件里单词出现的频率。 分支循坏
//石头剪刀布
//{
//系统产生1-3的随机数，分别代表剪刀，石头和布
// 玩家键盘输入1-3数字，分别代表剪刀，石头和布
//       }

func Test_Goto_11(t *testing.T) {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN>]]\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	frequencyForWord := map[string]int{} // 与:make(map[string]int)相同
	for _, filename := range commandLineFiles(os.Args[1:]) {
		updateFrequencies(filename, frequencyForWord)
	}
	reportByWords(frequencyForWord)
	wordsForFrequency := invertStringIntMap(frequencyForWord)
	reportByFrequency(wordsForFrequency)
}
func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // 无效模式
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}
func updateFrequencies(filename string, frequencyForWord map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer file.Close()
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}
func readAndUpdateFrequencies(reader *bufio.Reader,
	frequencyForWord map[string]int) {
	for {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax ||
				utf8.RuneCountInString(word) > 1 {
				frequencyForWord[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}
func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}
func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])
	}
}
func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
