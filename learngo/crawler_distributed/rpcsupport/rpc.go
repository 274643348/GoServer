package rpcsupport

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)


//注册一个service到host
func ServeRpc(host string,service interface{}) error {
	rpc.Register(service)

	listener,err:=net.Listen("tcp",host)
	if err != nil {
		return err
	}

	for {
		conn,err:=listener.Accept()
		if err != nil {
			fmt.Printf("accept error :%v",err)
		}
		fmt.Printf("accetp %v",conn)
		jsonrpc.ServeConn(conn)
	}
	return nil

}


//创建一个host的rpc对象，通过Call来调用
func NewClient(host string)(*rpc.Client,error){
	conn,err :=net.Dial("tcp",host)
	if err != nil {
		return nil,err
	}

	return jsonrpc.NewClient(conn),nil

}