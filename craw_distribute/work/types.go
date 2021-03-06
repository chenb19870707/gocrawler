package work

import (
	"gocrawler/engine"
	"gocrawler/parse/zhenai"
	"log"
	"fmt"
	"errors"
)

type SerializeParser struct{
	Name string
	Args interface{}
}

type Request struct{
	Url string
	Parse SerializeParser
}

type ParseResult struct{
	Items [] interface{}
	Requests []Request
}

func SerializeResult(r engine.ParseResult) ParseResult{
	result:=ParseResult{Items:r.Items}

	for _,req := range r.Requests{
		result.Requests = append(result.Requests,SerializeRequest(req))
	}
	return result
}
func SerializeRequest(r engine.Request) Request {

	name,args:= r.Parse.Serialize()

	return Request{
		Url:r.Url,
		Parse:SerializeParser{
			Name:name,
			Args:args,
		},
	}
}


func DeserializeResult(r ParseResult) engine.ParseResult{
	result:= engine.ParseResult{
		Items:r.Items,
	}

	for _,req := range r.Requests{
		engineReq,err:= DeserializeRequest(req)
		if err!=nil{
			log.Printf("error deserializeing:%v",err)
			continue
		}
		result.Requests = append(result.Requests,engineReq)
	}

	return result
}
func DeserializeRequest(r Request) (engine.Request, error) {
	parse,err:= deserializeParse(r.Parse)
	if err!=nil{
		return engine.Request{},err
	}
	return engine.Request{
		Url:r.Url,
		Parse:parse,
	},nil


}
func deserializeParse(p SerializeParser) (engine.Parse, error) {

	switch  p.Name {
	case "ParseCityList":
		return engine.NewFuncparse(zhenai.ParseCityList, "ParseCityList"), nil
	case "ParseCity":
		return engine.NewFuncparse(zhenai.ParseCity, "Parsecity"), nil

	case "ProfileParse":
		if useName, ok := p.Args.(string); ok {
			return zhenai.NewprofileParse(useName),nil
		}else{
			return nil,fmt.Errorf("invilid args:%v",p.Args)
		}
	case "Nilparse":
		return engine.NilParse{},nil
	default:
		return nil,errors.New("unkown parse name")

	}
}
