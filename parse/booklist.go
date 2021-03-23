package parse

import (
	"gocrawler/engine"
	"regexp"
)

const BooklistRe = `<a href="([^"]+)" title="([^"]+)"`
func ParseBookList(contents []byte)  engine.ParseResult{
	re := regexp.MustCompile(BooklistRe)

	matchs := re.FindAllSubmatch(contents,-1)

	result :=  engine.ParseResult{}

	for _,m :=range matchs{
		bookname := string(m[2])
		result.Items =append(result.Items,string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			string(m[1]),
			func(bytes []byte) engine.ParseResult {
				return ParseBookDetail(bytes,bookname)
			},
		})
	}

	return result
}