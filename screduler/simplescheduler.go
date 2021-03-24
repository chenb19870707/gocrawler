package screduler

import "gocrawler/engine"

type SimpleScheduler struct {
	wokerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request)  {
	go func() {
		s.wokerChan <- request
	}()
}

func (s *SimpleScheduler) Run()  {
	s.wokerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkChan() chan engine.Request{
	return s.wokerChan
}

func (s *SimpleScheduler) WorkReady(w chan engine.Request){
	return
}

