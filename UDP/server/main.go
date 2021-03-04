package main

import (
	"fmt"
	"net"
)

// UDP Server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 10002,
	})
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	defer listen.Close()

	for {
		var data [1024]byte

		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed,err:", err)
			return
		}
		fmt.Printf("data:%v addr:%v count:%v \n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
