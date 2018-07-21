package main

import (
	"net"
	"runtime"
	"sync"
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
	var wg sync.WaitGroup

	for i := 0; i < runtime.NumCPU()*1000;i++{
		wg.Add(1)
		data := fmt.Sprintf("hello-%d", i)
		fmt.Println(data)
		go func(i int){
			defer wg.Done()

			_ ,_,  err := conn.WriteMsgUDP([]byte(data),nil,nil)
			if err != nil{
				fmt.Println(err)
				return
			}

			//data := make([]byte, 1024)
			////n , err := conn.Read(data)
			//n , remoteAddr ,err := conn.ReadFromUDP(data)
			//if err != nil{
			//	fmt.Println(err)
			//	return
			//}
			////fmt.Printf("from [%s], recv [%s]\n", conn.RemoteAddr().String(), data[:n])
			//fmt.Printf("from [%s], recv [%s] id [%d]\n", remoteAddr.String(), data[:n],i)
		}(i)
	}
	wg.Wait()
}
