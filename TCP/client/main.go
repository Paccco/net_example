package main

import (
	"fmt"
	"net"
	"net_example/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10001")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := `paco `

		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err:", err)
			return
		}

		_, _ = conn.Write(data)
	}
}
