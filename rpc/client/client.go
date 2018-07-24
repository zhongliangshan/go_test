package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type GraphItem struct {
	Endpoint  string            `json:"endpoint"`
	Metric    string            `json:"metric"`
	Tags      map[string]string `json:"tags"`
	Value     float64           `json:"value"`
	Timestamp int64             `json:"timestamp"`
	DsType    string            `json:"dstype"`
	Step      int               `json:"step"`
	Heartbeat int               `json:"heartbeat"`
	Min       string            `json:"min"`
	Max       string            `json:"max"`
}

type SimpleRpcResponse struct {
	Code int `json:"code"`
}

type NullRpcRequest struct {}


func main() {
	conn, err := net.Dial("tcp", "180.97.84.130:9999")
	if err != nil {
		panic(err)
	}

	client := rpc.NewClient(conn)

	var result SimpleRpcResponse
	var request []GraphItem

	request = append(request , GraphItem{
		Endpoint : "tw13c962",
		Metric : "falcon",
		Value : 1,
	})

	err = client.Call("Graph.Send", request, &result)
	if err != nil {
		fmt.Println("error :", err)
	} else {
		fmt.Println(result)
	}

}
