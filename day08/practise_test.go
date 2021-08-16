package day08

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

// 爬邮箱
func GetEmail() {
	// 1.去网站拿数据
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	pageStr := string(pageBytes)
	//fmt.Println(pageStr)
	// 3.过滤数据，过滤qq邮箱
	re := regexp.MustCompile(reQQEmail)
	// -1代表取全部
	results := re.FindAllStringSubmatch(pageStr, -1)
	//fmt.Println(results)

	// 遍历结果
	for _, result := range results {
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}
}

// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func Test_goroutine_practise_01(t *testing.T) {
	GetEmail()
}
