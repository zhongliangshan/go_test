package main

import (
	"github.com/zhongliangshan/test/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdemo.QuickRpcDemo{})
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
			log.Println("accept err : ", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
