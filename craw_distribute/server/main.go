package main

import (
	"gocrawler/craw_distribute/persist"
	"gocrawler/craw_distribute/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main()  {
	ServerRpc(":1234")
}

func ServerRpc(host string) error {
	client,err :=	elastic.NewClient(elastic.SetSniff(false))
	if err !=nil{
		return err
	}

	return rpcsupport.ServerRpc(host,&persist.ItemService{
		client,
	})
}
