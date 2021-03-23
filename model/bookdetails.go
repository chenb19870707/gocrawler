package model

import "strconv"

type Bookddetail struct {
	BookName string		//书籍名字
	Author string		//作者
	Publicer string		//出版社
	Bookpages int		//页数
	Price string		//价格
	Score string		//豆瓣评分
	Intro string		//简介
}

func (b Bookddetail) String() string {
	return "书籍名字:" + b.BookName + " 作者:" + b.Author + " 出版社:" + b.Publicer + " 书籍页数:" + strconv.Itoa(b.Bookpages) +" 评分：" + b.Score+ " 价格:" + b.Price + "\n简介：" + b.Intro
}