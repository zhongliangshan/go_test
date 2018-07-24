package httpTest

import (
	"net"
	"log"
)

func main() {
	rAddr := &net.UDPAddr{IP: net.ParseIP("239.255.255.250"), Port: 3000}
	if conn, e := net.DialUDP("udp", nil, rAddr); e == nil {
		defer conn.Close()
		if _, e := conn.Write([]byte("q313213")); e == nil {
			log.Println("发送广播成功。")
		} else {
			log.Println("发送广播失败：", e)
		}
	}
}
