package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}
type Res struct {
	Num int
}

func main() {
	req := Req{
		Num1: 1,
		Num2: 2,
	}
	var res Res
	// rpc调用
	client, _ := rpc.DialHTTP("tcp", "localhost:8848")
	client.Call("Server.Add", req, &res)
	fmt.Println(res.Num)
}
