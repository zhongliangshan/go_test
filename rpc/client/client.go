package main

import (
	"net"
	"fmt"
	"github.com/zhongliangshan/test/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result []int
	var res float64
	inputs := []int{20, 1, 21, 7, 20, 4, 77, 1, 22, 0}
	fmt.Println(inputs)
	err = client.Call("QuickRpcDemo.QuickSort",rpcdemo.Args{inputs} , &result)
	if err  != nil {
		fmt.Println("error :" , err)
	} else {
		fmt.Println(result)
	}
	err = client.Call("QuickRpcDemo.Div",rpcdemo.Args2{10,3} , &res)

	if err  != nil {
		fmt.Println("error :" , err)
	} else {
		fmt.Println(res)
	}


}
