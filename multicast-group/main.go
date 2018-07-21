package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	addr,err := net.ResolveUDPAddr("udp","224.0.0.255:7878")
	if err != nil{
		panic(err)
	}

	listener ,  err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil{
		panic(err)
	}

	listener2 ,  err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil{
		panic(err)
	}

	go func (){
		data := make([]byte, 1024)
		for{
			n , remoteAddr, err := listener.ReadFromUDP(data)
			if err !=nil{
				fmt.Println(err)
				continue
			}
			fmt.Printf("from listener1 %s,read [%s]\n", remoteAddr.String(), string(data[:n]))
		}
	}()
	go func (){
		data := make([]byte, 1024)
		for{
			n , remoteAddr, err := listener2.ReadFromUDP(data)
			if err !=nil{
				fmt.Println(err)
				continue
			}
			fmt.Printf("from listener2 %s,read [%s]\n", remoteAddr.String(), string(data[:n]))
		}
	}()

	data := make([]byte, 1)
	os.Stdin.Read(data)
}
