package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net_example/proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		//n, err := reader.Read(buf[:])
		msg, err := proto.Decode(reader)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}

		fmt.Println(" 收到client端发来的数据:", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10001")

	if err != nil {
		fmt.Println("listen failed err:", err)
		return
	}
	for {
		conn, err := listen.Accept() //建立链接
		if err != nil {
			fmt.Println("accept failed err:", err)
			continue
		}
		go process(conn)
	}
}
