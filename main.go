package main

import (
	"gocrawler/engine"
	"gocrawler/parse"
	"gocrawler/persist"
	"gocrawler/screduler"
)

func main()  {
	e := engine.ConcurrentEngine{
		&screduler.SimpleScheduler{},
		100,
		persist.ItemSave(),
	}

	e.Run(engine.Request{
		"https://book.douban.com/",
		parse.ParseTag,
		//"http://www.zhenai.com/zhenghun",
		//zhenai.ParseCity,
	})
}

