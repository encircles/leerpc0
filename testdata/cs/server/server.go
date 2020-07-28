package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.Accept()
		//defer closeConn(conn)
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func closeConn(conn net.Conn) {
	fmt.Println("defer close")
	conn.Close()
}

func handleConn(conn net.Conn) {
	defer closeConn(conn)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	msg := string(buffer[:n])
	fmt.Println("recv from client: ", msg)

	_, _ = conn.Write([]byte("world"))
}
