package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (s SimpleEngine)Run(seeds... Request)  {

	var requests []Request
	for _,e := range seeds{
		requests = append(requests,e)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetch url:%s",r.Url)
		body,err := fetcher.Fetch(r.Url)

		if err != nil{
			log.Printf("Fecth Error %s",r.Url)
		}
		parseResult :=	r.ParseFnc(body)

		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items{
			log.Printf("Got item %s ",item)
		}
	}
}
