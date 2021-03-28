package parse

import (
	"gocrawler/engine"
	"regexp"
)

const regexStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`
func ParseTag(content []byte,_ string)  engine.ParseResult{
	re,_ := regexp.Compile(regexStr)

	match := re.FindAllSubmatch(content,-1)

	result := engine.ParseResult{}

	for _,m := range match{
		result.Items = append(result.Items,m[2])
		result.Requests = append(result.Requests,engine.Request{
			"https://book.douban.com" + string(m[1]),
			engine.NewFuncparse(ParseTag,"ParseTag"),
		})
	}

	//for _,m := range match{
	//	fmt.Printf("url:%s\n","https://book.douban.com" + string(m[1]))
	//}

	return result
}
