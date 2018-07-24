package common

import (
	"fmt"
	"net"
	"bytes"
)
const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:9000"
	DELIMITER = "\t"


)
var StartServer = make(chan int)
var logSn  =1
func PrintLog(msg string , args ...interface{}) {
	fmt.Printf("%d : %s" , logSn , fmt.Sprintf(msg , args...))
	logSn ++
}


func Read(conn net.Conn) (string , error) {
	var buffer bytes.Buffer

	readBytes := make([]byte , 1)

	for {
		_, err := conn.Read(readBytes)
		if err!= nil {
			return "" ,err
		}

		readByte := readBytes[0]

		if string(readByte) == DELIMITER {
			break
		}

		buffer.WriteByte(readByte)
	}
	return buffer.String() , nil
}

func Write(conn net.Conn , str string)(n int, err error) {
	var buffer bytes.Buffer
	buffer.WriteString(str)
	buffer.WriteString(DELIMITER)

	return conn.Write(buffer.Bytes())
}
