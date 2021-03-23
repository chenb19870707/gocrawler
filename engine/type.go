package engine

type ParseResult struct {
	Requests []Request
	Items[] interface{}
}

type Request struct {
	Url string
	ParseFnc func([] byte) ParseResult
}

func NilParse([] byte) ParseResult {
	return ParseResult{}
}


type Scheduler interface {
	Submit(request Request)
	ConfigureWorkChan(chan Request)
}
