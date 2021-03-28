package main

import (
	"gocrawler/craw_distribute/client"
	"gocrawler/engine"
	"gocrawler/parse/zhenai"
	"gocrawler/screduler"
)

func main()  {

	itemsave,err:= client.ItemSave(":1234")
	if(err != nil){
		panic(err)
	}

	e := engine.ConcurrentEngine{
		&screduler.SimpleScheduler{},
		100,
		itemsave,
	}

	e.Run(engine.Request{
		//"https://book.douban.com/",
		//parse.ParseTag,
		"http://www.zhenai.com/zhenghun",
		zhenai.ParseCity,
	})
}
