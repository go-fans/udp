package main

import (
	"net"
	"fmt"
)

/**
    func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
*/

func main() {
	srcAddr := &net.UDPAddr{IP:net.IPv4zero, Port:0}
	destAddr := &net.UDPAddr{IP:net.ParseIP("127.0.0.1"), Port:7878}

	conn , err := net.DialUDP("udp",srcAddr, destAddr)	//connected
	if err != nil{
		panic(err)
	}
	defer conn.Close()

	conn.WriteMsgUDP([]byte("hello"),nil,nil)

	//conn.Write([]byte("hello"))

	data := make([]byte, 1024)
	//n , err := conn.Read(data)
	n , remoteAddr ,err := conn.ReadFromUDP(data)
	if err != nil{
		fmt.Println(err)
		return
	}
	//fmt.Printf("from [%s], recv [%s]\n", conn.RemoteAddr().String(), data[:n])
	fmt.Printf("from [%s], recv [%s]\n", remoteAddr.String(), data[:n])
}
