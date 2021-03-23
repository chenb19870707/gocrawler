package parse

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
	"strings"
)

var authorRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var publicRe = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var pagesRe =  regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceRe =  regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br/>`)
var scoreRe =  regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var introRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseBookDetail(contents []byte,bookname string) engine.ParseResult {
	bookdetail := model.Bookddetail{}

	bookdetail.Author = extraString(contents,authorRe)
	pageStr:= strings.Trim(extraString(contents,pagesRe),"")
	page,err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}
	bookdetail.Bookpages = page
	bookdetail.Price = extraString(contents,priceRe)
	bookdetail.Intro = extraString(contents,introRe)
	bookdetail.Score = extraString(contents,scoreRe)
	bookdetail.Publicer = extraString(contents,publicRe)
	bookdetail.BookName = bookname

	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}
	return result
}

func extraString(contents []byte,re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >=  2{
		return string(match[1])
	}else{
		return ""
	}
}
