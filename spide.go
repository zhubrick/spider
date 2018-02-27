package main

import "fmt"
import "os"
import "io/ioutil"
import "regexp"
import "net/http"
import "log"
import iconv "github.com/djimenez/iconv-go"

type Spider struct {
	Page int
}

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

	out := make([]byte, len(data))
	out = out[:]

	iconv.Convert(data, out, "gb2312", "utf-8")
	content = string(out)
	statusCode = resp.StatusCode

	return

}

func (this *Spider) Spider_one_page() {
	fmt.Println("正在爬取", this.Page, "页")
	utlhead := "https://www.neihanba.com"
	var res string
	var stat int
	if this.Page == 1 {
		url := utlhead + "/dz/index.html"
		res, stat = this.HttpGet(url)
	} else {
		url := fmt.Sprintf("%s/dz/list_%d.html", utlhead, this.Page)
		res, stat = this.HttpGet(url)
	}
	if stat == 200 {
		dirname := fmt.Sprintf("%d", this.Page)
		err := os.MkdirAll(dirname, 0777)
		if err != nil {
			return
		}
		myreg := regexp.MustCompile(`<h4> <a href="(.*?)"`)
		tail := myreg.FindAllStringSubmatch(res, -1)

		for i := 0; i < len(tail); i++ {
			url := utlhead + tail[i][1]
			fmt.Println(url)
			res, _ = this.HttpGet(url)
			filename := fmt.Sprintf("./%s/page%d-%d.txt", dirname, this.Page, i+1)

			f, ferr := os.Create(filename)
			if ferr != nil {
				fmt.Println(ferr)
				return
			}

			titlemyreg := regexp.MustCompile(`<h1>(.*?)</h1>`)
			title := titlemyreg.FindAllStringSubmatch(res, -1)
			f.WriteString(title[0][1])
			f.WriteString("\n")
			bodymyreg := regexp.MustCompile(`<td><p>(?s:(.*?))</p></td>`)
			body := bodymyreg.FindAllStringSubmatch(res, -1)

			f.WriteString(body[0][1])
			f.Close()
		}
	}
}

func (this *Spider) Work() {
	this.Page = 1
	var cmd string

	for {
		fmt.Println("请输入任意键爬取下一页，输入exit退出")
		fmt.Scanf("%s", &cmd)
		if cmd == "exit" {
			fmt.Println("exit")
			break
		}
		this.Spider_one_page()
		this.Page++
	}

}

func main() {
	p := new(Spider)
	p.Work()

}
