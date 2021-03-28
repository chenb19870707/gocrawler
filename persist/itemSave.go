package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSave()  chan interface{}{
	out := make(chan interface{})

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Printf("es error")
	}

	go func() {
		itemcount := 0
		for  {
			item := <- out
			log.Printf("Item saver:Got$%d,$v",itemcount,item)
			Save(client,item)
			itemcount++
		}
	}()

	return out
}

func Save(client *elastic.Client,item interface{}) error{
	_,err := client.Index().Index("dating_profile").Type("zhihu").BodyJson(item).Do(context.Background())
	if err!=nil{
		return err
	}
	return nil
}
