package main

import (
	"github.com/zhongliangshan/test/httpTest/http/server"
	"github.com/zhongliangshan/test/httpTest/http/common"
	"github.com/zhongliangshan/test/httpTest/http/client"
)

func main() {
	go server.StartServer(common.SERVER_NETWORK , common.SERVER_ADDRESS)

	done := <-common.StartServer

	common.PrintLog("start server %d\n" , done)

	go clients.StartClient(common.SERVER_NETWORK , common.SERVER_ADDRESS)
	done2 := <-common.StartServer
	common.PrintLog("start client %d\n" , done2)
}
