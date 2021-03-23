package main

import (
	"crawler/engine"
	"crawler/parse"
	"crawler/screduler"
)

func main()  {
	e := engine.ConcurrentEngine{
		&screduler.SimpleScheduler{},
		100,
	}

	e.Run(engine.Request{
		"https://book.douban.com/",
		parse.ParseTag,
	})
}

