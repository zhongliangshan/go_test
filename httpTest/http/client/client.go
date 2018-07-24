package clients

import (
	"net"
	"time"
	"github.com/zhongliangshan/test/httpTest/http/common"
	"math/rand"
	"fmt"
)

func StartClient(network  , address string) {
	conn, err := net.DialTimeout(network, address, 2*time.Second)
	if err != nil {
		common.PrintLog("connection server error:%s\n" , err)
		return
	}

	defer conn.Close()

	common.PrintLog("Connected to server. (remote address: %s, local address: %s) (Client[%d])\n",
		conn.RemoteAddr(), conn.LocalAddr())
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))

	// 发送数据
	for i:=0;i<5;i++ {
		request := rand.Int63()
		n, err := common.Write(conn, fmt.Sprintf("%d", request))

		if err != nil {
			common.PrintLog("Write Error: %s (Client[%d])\n", err)
			continue
		}
		common.PrintLog("Sent request (written %d bytes): %d (Client[%d])\n", n, request)

	}

	// 读取数据
	for i:=0;i<5;i++ {
		str, err := common.Read(conn)

		if err != nil {
			common.PrintLog("Read Error: %s (Client[%d])\n", err)
			continue
		}
		common.PrintLog("Read response :%s\n" , str)

	}

	common.StartServer <- 2
}


