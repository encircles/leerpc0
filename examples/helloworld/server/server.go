package main

import (
	"time"

	"github.com/encircles/leerpc0"
	"github.com/encircles/leerpc0/testdata/helloworld"
)

func main() {
	opts := []leerpc0.ServerOption{
		leerpc0.WithAddress("127.0.0.1:8000"),
		leerpc0.WithNetwork("tcp"),
		leerpc0.WithSerializationType("msgpack"),
		leerpc0.WithTimeout(time.Millisecond * 2000),
	}

	s := leerpc0.NewServer(opts...)
	if err := s.RegisterService("/helloworld.Greeter", new(helloworld.Service)); err != nil {
		panic(err)
	}

	s.Serve()
}
