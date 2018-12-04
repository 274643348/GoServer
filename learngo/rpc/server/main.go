package main

import (
	"fmt"
	"learngo/GoServer/learngo/rpc"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
	rpc.go中实现一个Div方法，然后在main中开启网络服务器
	在终端中输入 telnet localhost 1234
	然后在输入测试代码：{"method":"abc.def"}

	//测试：
	//{"method":"abc.def"}--错误演示
	//{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1314}--正确演示
 */
func main() {

	rpc.Register(rpcDemo.DemoService{})

	listener,err:=net.Listen("tcp",":1234");
	if err != nil {
		panic(err);
	}

	for {
		conn,err:=listener.Accept()
		if err != nil {
			fmt.Printf("accept error :%v",err)
		}
		fmt.Printf("accetp %v",conn)
		jsonrpc.ServeConn(conn)
	}



}
