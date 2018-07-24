package httpTest

import (
	"net"
	"log"
	"fmt"
)

func main() {
	udpAddr := net.UDPAddr{IP: net.ParseIP("239.255.255.250"), Port: 3000}
	udpConn, err := net.ListenMulticastUDP("udp", nil, &udpAddr)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	for {
		n, addr, err := udpConn.ReadFrom(buffer)
		if err != nil {
			log.Println("接收数据失败")
			continue
		}

		fmt.Println(buffer[:n] , addr)
	}

}
