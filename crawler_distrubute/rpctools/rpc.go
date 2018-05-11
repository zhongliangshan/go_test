package rpctools

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

type RpcParams struct {
	NetWork string
	Host string
}

type RpcServerParams struct {
	RpcParams RpcParams
	Service interface{}
}

func (rpcServer RpcServerParams) ServerRpc() error{
	err := rpc.Register(rpcServer.Service)
	if err != nil {
		return err
	}

	listen, err := net.Listen(rpcServer.RpcParams.NetWork , rpcServer.RpcParams.Host)
	if err != nil {
		return err
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("accept err : " , err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

	return nil
}

func (rpcClient RpcParams) ClientRpc() (*rpc.Client , error){
	conn, err := net.Dial(rpcClient.NetWork , rpcClient.Host)
	if err != nil {
		return nil , err
	}

	return jsonrpc.NewClient(conn) , nil
}