package main

import (
	"gocrawler/engine"
	"gocrawler/parse"
	"gocrawler/screduler"
)

func main()  {
	e := engine.ConcurrentEngine{
		&screduler.QueueScheduler{},
		100,
	}

	e.Run(engine.Request{
		"https://book.douban.com/",
		parse.ParseTag,
	})
}

