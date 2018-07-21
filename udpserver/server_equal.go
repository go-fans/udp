package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

func read(conn *net.UDPConn){
	for{
		data := make([]byte,1024)
		//n , err := conn.Read(data)
		n, remoteAddr , err := conn.ReadFromUDP(data) //Read ??
		if err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Printf("From [%s] , recv [%s]\n", remoteAddr.String(),  string(data[:n]))
	}
}

// unconnected
func serverEqual(){
	addr1 := &net.UDPAddr{IP:net.ParseIP("127.0.0.1"), Port:7879}
	addr2 := &net.UDPAddr{IP:net.ParseIP("127.0.0.1"), Port:7878}

	//l1
	go func(){
		listener, err := net.ListenUDP("udp", addr1)
		if err != nil{
			panic(err)
		}
		go read(listener)
		time.Sleep(time.Second*1)
		//listener.WriteToUDP([]byte("ping 2"), addr2)
		listener.WriteMsgUDP([]byte("ping 2"),nil,  addr2)
	}()

	//l2
	go func(){
		listener, err := net.ListenUDP("udp", addr2)
		if err != nil{
			panic(err)
		}
		go read(listener)
		time.Sleep(time.Second*1)
		listener.WriteMsgUDP([]byte("ping 1"),nil,  addr1)
		//listener.WriteToUDP([]byte("ping 1"), addr1)
	}()

	b := make([]byte,1)
	os.Stdin.Read(b)	//pause ...
}

