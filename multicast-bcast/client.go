package main

import (
	"net"
	"fmt"
)

func main() {
	srcAddr := &net.UDPAddr{IP:net.IPv4zero, Port:0}
	destAddr := &net.UDPAddr{IP:net.ParseIP("192.168.47.255"), Port:7878}

	conn , err := net.ListenUDP("udp",srcAddr)	//connected
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	conn.WriteMsgUDP([]byte("hello"),nil,destAddr)
	data := make([]byte, 1024)
	n , remoteAddr , err  := conn.ReadFromUDP(data)
	if err !=nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("From [%s], recv [%s]\n", remoteAddr.String(), string(data[:n]))
}
