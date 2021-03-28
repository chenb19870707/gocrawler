package main

import (
	"gocrawler/craw_distribute/rpcsupport"
	"gocrawler/craw_distribute/work/server"
	"log"
)

func main()  {
	log.Fatal(rpcsupport.ServerRpc(":1235",&server.CrawlService{}))
}

