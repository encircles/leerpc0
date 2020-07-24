package main

import (
	"time"

	"github.com/encircles/leerpc0"
)

func main() {
	opts := []leerpc0.ServerOption{
		leerpc0.WithAddress("127.0.0.1:8000"),
		leerpc0.WithNetwork("tcp"),
		leerpc0.WithProtocol("proto"),
		leerpc0.WithTimeout(time.Millisecond * 2000),
	}

	_ = leerpc0.NewServer(opts...)

}
