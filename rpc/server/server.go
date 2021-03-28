package main

import (
	rpcdemo "gocrawler/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	rpc.Register(rpcdemo.DemoService{})

	listen,err := net.Listen("tcp",":1234")
	if err!=nil{
		panic(err)
	}

	for  {
		conn,err := listen.Accept()
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}