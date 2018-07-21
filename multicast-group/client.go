package main

import (
	"net"
)

func main() {
	srcAddr := &net.UDPAddr{IP:net.IPv4zero, Port:0}
	destAddr := &net.UDPAddr{IP:net.ParseIP("224.0.0.255"), Port:7878}

	conn , err := net.DialUDP("udp",srcAddr, destAddr)	//connected
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	conn.WriteMsgUDP([]byte("hello"),nil,nil)
}
