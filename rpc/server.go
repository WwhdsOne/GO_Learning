package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct{}
type Req struct {
	Num1 int
	Num2 int
}
type Res struct {
	Num int
}

func (s *Server) Add(req Req, res *Res) error {
	fmt.Printf("%v", req)
	res.Num = req.Num1 + req.Num2
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8848")
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Serve(l, nil)
}
