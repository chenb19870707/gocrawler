package client

import (
	"gocrawler/craw_distribute/rpcsupport"
	"log"
)

func ItemSave(host string)  (chan interface{},error){
	out := make(chan interface{})

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil,err
	}

	go func() {
		itemcount := 0
		for  {
			item := <- out
			log.Printf("Item saver:Got$%d,$v",itemcount,item)
			result := ""
			err := client.Call("ItemService.Save",item,&result)
			if err != nil{
				log.Printf("item save error saving item %v%v",item,err)
			}
			itemcount++
		}
	}()

	return out,nil
}
