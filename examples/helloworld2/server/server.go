package main

import (
	"context"
	"fmt"
	"time"

	"github.com/encircles/leerpc0"
	"github.com/encircles/leerpc0/examples/helloworld2/helloworld"
)

type greeterService struct{}

func (g *greeterService) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println("recv Msg : ", req.Msg)
	rsp := &helloworld.HelloReply{
		Msg: req.Msg + " world",
	}
	return rsp, nil
}

func main() {
	opts := []leerpc0.ServerOption{
		leerpc0.WithAddress("127.0.0.1:8000"),
		leerpc0.WithNetwork("tcp"),
		leerpc0.WithProtocol("proto"),
		leerpc0.WithTimeout(time.Millisecond * 2000),
	}

	s := leerpc0.NewServer(opts...)

	helloworld.RegisterService(s, &greeterService{})

	s.Serve()
}
