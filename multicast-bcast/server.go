package main

import (
	"net"
	"fmt"
)

func main() {
	listener ,err := net.ListenUDP("udp",&net.UDPAddr{IP:net.IPv4zero, Port:7878})
	if err != nil{
		panic(err)
	}

	data := make([]byte, 1024)
	for{
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Printf("From [%s], recv [%s]\n", remoteAddr.String(), string(data[:n]))
		_ , _ , err = listener.WriteMsgUDP([]byte("world"),nil,remoteAddr)
		if err != nil{
			fmt.Println(err)
			continue
		}
	}
}
