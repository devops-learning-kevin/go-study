package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

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

func (s *Server) Add(req Req, res *Res) error {
	res.Num = req.NumOne + req.NumTow
	return nil
}

func (s *Server) Sub(req Req, res *Res) error {
	res.Num = req.NumOne - req.NumTow
	return nil
}

func (s *Server) GetKey(Keyreq KeyReq, keyres *KeyRes) error {
	keyres.Address = "0xD9AB964174EF396806EB2B86BCD05E59F605C3AC"
	keyres.Private = "18787fbd765e17e9eac6390a7bbadbeccdabbf84790ff105ff30a4b801d1cafe"
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Serve(l, nil)
}
