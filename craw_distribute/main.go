package main

import (
	"gocrawler/craw_distribute/client"
	"gocrawler/engine"
	client2 "gocrawler/craw_distribute/work/client"
	"gocrawler/parse/zhenai"
	"gocrawler/screduler"

)

func main(){

	itemsave,err:= client.ItemSave(":1234")

	process,err:=  client2.CreateProcess()
	if err!=nil{
		panic(err)
	}
	e:= engine.ConcurrentEngine{
		&screduler.QueueScheduler{},
		100,
		itemsave,
		process,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		Parse:engine.NewFuncparse(zhenai.ParseCityList,"ParseCityList") ,
	})
}