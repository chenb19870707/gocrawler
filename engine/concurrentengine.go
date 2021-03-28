package engine

import (
	"gocrawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

func (e *ConcurrentEngine) Run(seeds...Request)  {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0;i<e.WorkerCount;i++{
		CreateWoker(e.Scheduler.WorkChan(),out,e.Scheduler)
	}

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	for{
		result := <-out
		for _,item := range result.Items{
			go func() {
				e.ItemChan <- item
			}()
		}

		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func CreateWoker(in chan Request,out chan ParseResult,s Scheduler)  {
	go func() {
		for  {
			s.WorkReady(in)
			request := <-in

			result,err := woker(request)
			if err != nil{
				continue
			}

			out <- result
		}
	}()
}

func woker(request Request) (ParseResult, error) {
	//fmt.Printf("Fecth url %s\n",request.Url)

	body,err := fetcher.ProxyFetch(request.Url)
	if err!=nil{
		log.Printf("feth error:%s",err)
		return ParseResult{},err
	}

	return request.ParseFnc(body),nil
}