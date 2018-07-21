package main

import (
	"net"
	"fmt"
)

//main: create udp server

func main(){
	serverEqual()
	return
	listener , err := net.ListenUDP("udp",&net.UDPAddr{IP:net.ParseIP("127.0.0.1"),Port:7878})
	if err != nil{
		panic(err)
	}
	fmt.Printf("Local <%s>\n", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for{
		n , remoteAddr, err := listener.ReadFromUDP(data)
		if err !=nil{
			fmt.Println(err)
			continue
		}
		fmt.Printf("from %s,read [%s]\n", remoteAddr.String(), string(data[:n]))
		_ ,err = listener.WriteToUDP([]byte("ok"), remoteAddr)
		if err !=nil{
			fmt.Println(err)
			continue
		}
	}
}
