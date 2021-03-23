package engine

import (
	"crawler/fetcher"
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(seeds...Request)  {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfigureWorkChan(in)

	for i := 0;i<e.WorkerCount;i++{
		CreateWoker(in,out)
	}

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount := 0;
	for{
		result := <-out
		for _,item := range result.Items{
			fmt.Printf("Got Item:%d,%v\n",itemCount,item)
			itemCount++
		}

		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func CreateWoker(in chan Request,out chan ParseResult)  {
	go func() {
		for  {
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
	fmt.Printf("Fecth url %s\n",request.Url)

	body,err := fetcher.Fetch(request.Url)
	if err!=nil{
		log.Printf("feth error:%s",err)
		return ParseResult{},err
	}

	return request.ParseFnc(body),nil
}