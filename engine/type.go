package engine

type ParseResult struct {
	Requests []Request
	Items[] interface{}
}

type Parse interface {
	Parse(contents []byte,url string) ParseResult
	Serialize()(name string,args interface{})
}

type Request struct {
	Url string
	Parse Parse
}


type NilParse struct {
	
}

func (n NilParse) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (n NilParse) Serialize() (name string, args interface{}) {
	return "NilParse",nil
}

type Scheduler interface {
	Submit(request Request)
	Run()
	WorkReady(w chan Request)
	WorkChan() chan Request
}

type ParseFunc func(contents []byte,url string) ParseResult

type FuncParses struct {
	parser ParseFunc
	name string
}

func (f FuncParses) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents,url)
}

func (f FuncParses) Serialize() (name string, args interface{}) {
	return f.name,nil
}

func NewFuncparse(p ParseFunc,name string) *FuncParses{
	return &FuncParses{
		parser:p,
		name:name,
	}
}

type Item struct{

	Url string
	Type string
	Id string
	Payload interface{}
}