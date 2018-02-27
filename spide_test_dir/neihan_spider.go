package main

import "fmt"
import "net/http"
import "log"
import "io/ioutil"

//明确目标
//第1页  https://www.neihanba.com/dz/index.html
//第2页  https://www.neihanba.com/dz/list_2.html
//第n页 https://www.neihanba.com/dz/list_n.html

//1 首先进入某页的页码主页，----> 取出每个段子链接地址
// https://www.neihanba.com 拼接一个段子的完整url路径
//得到每个段子路径的正则表达式  `<h4> <a href="(?:s(.*?))"`

// https://www.neihanba.com + /dz/1092886.html

// 进入每个段子的首页，得到段子的标题和内容

//标题的正则
//`<h1>(?:s(.*?))</h1>`

//内容的正则
//`<td><p>(?s:(.*?))</p></td>`

type Spider struct {
	Page int //当前爬虫已经爬取到了第几页
}

//爬取一个某页的菜单页码
func (this *Spider) Spider_one_page() {
	fmt.Println("正在爬取 ", this.Page, " 页")

}

//请求一个页码将页码中的全部源码content
func (this *Spider) HttpGet(url string) (content string, statusCode int) {

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		content = ""
		statusCode = -100
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		statusCode = resp.StatusCode
		content = ""
		return
	}

	content = string(data)
	statusCode = resp.StatusCode

	return
}

func (this *Spider) DoWork() {

	fmt.Println("Spider begin to  work")
	this.Page = 1

	var cmd string

	for {
		fmt.Println("请输入任意键爬取下一页，输入exit退出")
		fmt.Scanf("%s", &cmd)
		if cmd == "exit" {
			fmt.Println("exit")
			break
		}

		//需要爬取下一页
		this.Spider_one_page()

		this.Page++

	}

}

func main() {

	sp := new(Spider)
	sp.DoWork()

}
