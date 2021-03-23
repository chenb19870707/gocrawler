package screduler

import "crawler/engine"

type SimpleScheduler struct {
	wokerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request)  {
	go func() {
		s.wokerChan <- request
	}()
}

func (s *SimpleScheduler) ConfigureWorkChan(c chan engine.Request)  {
	s.wokerChan = c
}