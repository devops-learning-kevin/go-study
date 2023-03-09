package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	NumOne int
	NumTow int
}

type Res struct {
	Num int
}

type KeyReq struct {
}

type KeyRes struct {
	Address string
	Private string
}

func main() {
	req := Req{
		NumOne: 8,
		NumTow: 16,
	}
	var res Res
	var keyres KeyRes
	var keyreq KeyReq
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 同步调用
	client.Call("Server.Add", req, &res)
	fmt.Println(res)

	client.Call("Server.Sub", req, &res)
	fmt.Println(res)

	client.Call("Server.GetKey", keyreq, &keyres)
	fmt.Println(keyres)

	//异步调用
	//ca := client.Go("Server.Add", req, &res, nil)
	//for {
	//	select {
	//	case <-ca.Done:
	//		fmt.Println(res)
	//		return
	//	default:
	//		fmt.Println("处理别的任务")
	//	}
	//}

}
