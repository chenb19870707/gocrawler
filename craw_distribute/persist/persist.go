package persist

import (
	"gocrawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemService struct {
	Client *elastic.Client
}

func (s *ItemService) Save(item interface{},result *string)  error{
	err := persist.Save(s.Client,item)
	if err == nil{
		*result = "ok"
	}

	return err
}
