package engine

import (
	"gocrawler/fetcher"
	"log"
)


type Processor func( Request) (ParseResult,error)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
	RequestProcessor Processor
}

func (e *ConcurrentEngine) Run(seeds...Request)  {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0;i<e.WorkerCount;i++{
		e.CreateWoker(e.Scheduler.WorkChan(),out,e.Scheduler)
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

func (e *ConcurrentEngine)CreateWoker(in chan Request,out chan ParseResult,s Scheduler)  {
	go func() {
		for  {
			s.WorkReady(in)
			request := <-in

			result,err := e.RequestProcessor(request)
			if err != nil{
				continue
			}

			out <- result
		}
	}()
}

func Worker(request Request) (ParseResult, error) {
	//fmt.Printf("Fecth url %s\n",request.Url)

	body,err := fetcher.ProxyFetch(request.Url)
	if err!=nil{
		log.Printf("feth error:%s",err)
		return ParseResult{},err
	}

	return request.Parse.Parse(body,request.Url),nil
}