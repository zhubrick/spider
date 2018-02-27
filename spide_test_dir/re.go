package main

import "fmt"
import "regexp"

var content string = "3.14 123123 .68 haha 1.0 abc 7.7666 123."

var content2 string = `
<title>标题</title>
<div>xixi</div>
<div>haha</div>
<div>hei\n
6666
=55kai
3jijia3jiekui
hei</div>
<div>lala</div>
<body>bodyaasd</body>
`

func main() {

	//通过正则表达式 创建一个正则的句柄规则
	myreg := regexp.MustCompile(`\d+\.\d+`)

	//通过正则句柄 匹配源字符串， result是匹配得到结果
	result := myreg.FindAllStringSubmatch(content, -1)

	fmt.Printf("%+v\n", result)

	//[[3.14] [1.0] [7.7666]]

	for _, mystr_slice := range result {
		fmt.Printf("%s\n", mystr_slice[0])
	}

	fmt.Println("---------")

	divExp := regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	div_result := divExp.FindAllStringSubmatch(content2, -1)

	fmt.Printf("%+v\n", div_result)

	/*

		[[<div>xixi</div> xixi] [<div>haha</div> haha] [<div>hei\n
		6666
		=55kai
		3jijia3jiekui
		hei</div> hei\n
		6666
		=55kai
		3jijia3jiekui
		hei] [<div>lala</div> lala]]

	*/
	for _, text := range div_result {
		fmt.Println("======")
		fmt.Println(text[0]) //<div>
		fmt.Println("======")
		fmt.Println(text[1]) //no <div>
	}
}
