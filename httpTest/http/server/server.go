package server

import (
	"net"
	"github.com/zhongliangshan/test/httpTest/http/common"
	"time"
	"io"
	"strconv"
	"math"
	"fmt"
)

func StartServer(network  , address string){
	listener, err := net.Listen(network, address)
	if err != nil {
		common.PrintLog("open lister error:%s\n" , err)
		return
	}

	defer listener.Close()
	common.PrintLog("Got listener for this server:%s\n" , listener.Addr())
	common.StartServer <- 1
	for {
		conn, err := listener.Accept()

		if err != nil {
			common.PrintLog("accept error:%s\n" , err)
			continue
		}
		common.PrintLog("Got listener for remote server:%s\n" , conn.RemoteAddr())
		go handleConn(conn)


	}

}

func handleConn(conn net.Conn) {
	for  {
		// 设置连接的读取时间
		conn.SetDeadline(time.Now().Add(10*time.Second))

		str, err := common.Read(conn)
		if err != nil {
			if err == io.EOF {
				common.PrintLog("the connection close by another side\n")
			} else {
				common.PrintLog("read error:%s\n" , err)
			}

			break
		}

		common.PrintLog("read string:%s\n" , str)

		// 转成整型数据

		num , err :=convertInt64(str)
		if err != nil {

			_, e := common.Write(conn, err.Error())
			if e != nil {
				common.PrintLog("convertInt64 write conn error:%s\n" , e)
			}

			common.PrintLog("convertInt64 string error:%s\n" , err)
			continue
		}

		resp64 := math.Cbrt(float64(num))

		respMsg := fmt.Sprintf("The cube root of %d is %f.", num, resp64)

		_, e := common.Write(conn, respMsg)
		if e != nil {
			common.PrintLog("Cbrt write conn error:%s\n" , e)
		}
		common.PrintLog("Sent response (written %d bytes): %s (Server)\n",  respMsg)
	}
}

func convertInt64(str string) (int64 , error) {
	i , e := strconv.Atoi(str)
	if e != nil {
		return -1,e
	}

	if i > math.MaxInt64 || i < math.MinInt64 {
		return -1 , fmt.Errorf("str out number 64 bit")
	}

	return int64(i) , nil
}

