package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string,service interface{}) error {
	rpc.Register(service)
	listen,err := net.Listen("tcp",host)
	if err!=nil{
		return err
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

func NewClient(host string) (*rpc.Client,error) {
	conn,err := net.Dial("tcp",host)

	if err != nil{
		return nil,err
	}

	return jsonrpc.NewClient(conn),nil
}