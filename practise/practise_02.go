package practise

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
	"unicode"
	"unicode/utf8"
)

func Practise02() {
	//用来保存从文件读到的每一个单词和对应频率
	frequencyForWord := map[string]int{} // 与:make(map[string]int)相同
	//将命令行的文件名传入commandLineFiles函数中，并遍历取出文件名
	for _, filename := range commandLineFiles(os.Args[1:]) {
		//分析完文件后更新frequencyForWord的数据
		updateFrequencies(filename, frequencyForWord)
	}
	//输出第一个报告：按照字母顺序排序的单词列表和对应的频率
	reportByWords(frequencyForWord)
	//创建一个反转的映射map[int][] string，也就是说，键是频率而值则是所有具有这个频率的单词。并保存到wordsForFrequency
	wordsForFrequency := invertStringIntMap(frequencyForWord)
	//输出：按照出现频率排序的列表
	reportByFrequency(wordsForFrequency)
}

//因为Unix 类系统（如 Linux 或 Mac OS X 等）的 shell 默认会自动处理通配符，而Windows 平台的 shell 程序（cmd.exe）不支持通配符
//为了保持平台之间的一致性，这里使用 commandLineFiles() 函数来实现跨平台的处理
func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		//定义一个string类型的切片
		args := make([]string, 0, len(files))
		//将文件遍历取出有效值并添加到args切片中
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name) // //有err，无效模式
				// matches不为零值时
			} else if matches != nil {
				//将matches解好后添加到切片args中
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}

//处理文件
func updateFrequencies(filename string, frequencyForWord map[string]int) {
	var file *os.File
	var err error
	//打开文件失败
	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer file.Close()
	//将文件作为一个 *bufio.Reader传给 readAndUpdateFrequencies() 函数
	readAndUpdateFrequencies(bufio.NewReader(file), frequencyForWord)
}

//读取并更新frequencyForWord
func readAndUpdateFrequencies(reader *bufio.Reader,
	frequencyForWord map[string]int) {

	//无限循环来读取文件
	for {
		//一行行来读
		line, err := reader.ReadString('\n')
		//SplitOnNonLetters() 函数忽略掉非单词的字符，并且过滤掉字符串开头和结尾的空白
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			//len(word)>utf8.UTFMax：用来检査这个单词的字节数是否大于 utf8.UTFMax（常量，值为4）
			//utf8.RuneCountInString(word)>1：只记录含有两个以上（包括两个）字母的单词
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

//切分非单词字符的字符串（即空格，符号）
func SplitOnNonLetters(s string) []string {
	//匿名函数：如果传入的是字符那就返回 false，否则返回 true
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	//strings.FieldsFunc() 用于在每次运行的满足f(c)的Unicode代码点c处拆分给定的字符串str，并返回由str组成的切片数
	return strings.FieldsFunc(s, notALetter)
}

//反转，遍历原来的映射，将它的值作为键保存到反转的映射里，并将键增加到对应的值里去
func invertStringIntMap(intForString map[string]int) map[int][]string {
	//创建一个长度为原数据长度且为map[int]string
	stringsForInt := make(map[int][]string, len(intForString))
	//遍历原数据intForString
	for key, value := range intForString {
		//将原数据stringsForint的key保存到反转数据stringsForint的value
		//新的映射的值就是一个字符串切片，即使原来的映射有多个键对应同一个值，也不会丢掉任何数据。
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

//按照字母顺序排序的单词列表和对应的频率map[string]int输出
func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		//将单词追加到words切片中
		words = append(words, word)
		//utf8.RuneCountInString()：获取字符串中的字符个数
		//计算单词的宽度
		if width := utf8.RuneCountInString(word); width > wordWidth {
			//每次循环width就会和wordWidth比较，并更新wordWidth
			wordWidth = width
		}
		//fmt.Sprint(frequency)：使用默认格式为其操作数设置Sprint格式，并返回结果字符串，如果都不是字符串，则在操作数之间添加空格
		//同上这个是获取频率的宽度的
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	//给words切片排序即单词排序
	sort.Strings(words)
	//控制空格的个数
	//打印标题头：  单词Word+空格+频率Frequency,（%*s:" "）：中间隔*个空格
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	//循环遍历words只取值（即单词）并打印
	for _, word := range words {
		//%-*s  %*d\n  : 输出所有的字符型，输出所有的整型
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])
	}
}

//按照出现频率排序-单词的列表输出 wordsForFrequency :map[int][]string:key为int value为切片类型
func reportByFrequency(wordsForFrequency map[int][]string) {
	//创建一个切片用来保存频率
	frequencies := make([]int, 0, len(wordsForFrequency))
	//遍历词库，映射遍历key值
	for frequency := range wordsForFrequency {
		//将每一个单词频率添加到frequencies这个切片中
		frequencies = append(frequencies, frequency)
	}
	//按照升序排列频率
	sort.Ints(frequencies)
	//计算需要容纳的最大长度（频率最高的）并以此为第一列宽度
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		//取出每一个频率，并通过频率查找对应的单词
		words := wordsForFrequency[frequency]
		//按单词排序
		sort.Strings(words)
		//strings.Join(words,",")：用","号连接同一频率的单词
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}
