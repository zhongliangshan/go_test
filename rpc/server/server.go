package main

import (
	"github.com/zhongliangshan/test/rpc"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdemo.DemoRpc{})
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()

		if err != nil {
			continue
		}

		jsonrpc.ServeConn(conn)
	}

}
