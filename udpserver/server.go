package main

import (
	"net"
	"fmt"
	"runtime"
	"time"
)

//main: create udp server
var (
	recvCount = 0
	sendCount = 0
)

type Job struct{
	Conn *net.UDPConn
	Addr *net.UDPAddr
	Data string
}

var JobQueue = make(chan Job, 5000)

func worker(){
	for{
		select {
		case job := <-JobQueue:
			fmt.Printf("from %s,read [%s]\n", job.Addr.String(), job.Data)
			recvCount++
			//_ , _ ,err := job.Conn.WriteMsgUDP([]byte("ok"), nil, job.Addr)
			//if err !=nil{
			//	fmt.Println(err)
			//}
			default:
				time.Sleep(time.Second*2)
				fmt.Println(recvCount)
		}
	}
}

func listener(conn *net.UDPConn,id int, quit chan struct{}){
	data := make([]byte, 1024)
	fmt.Println("listener ",id)
	for{
		n , remoteAddr, err := conn.ReadFromUDP(data)
		if err !=nil{
			fmt.Println(err)
			continue
		}
		if len(JobQueue) == cap(JobQueue){
			fmt.Println("full...")
		}
		JobQueue <- Job{
			Conn:conn,
			Addr : remoteAddr,
			Data: string(data[:n]),
		}

	}
	fmt.Printf("listener %d over\n",id)
	quit <- struct{}{}
}

func main(){
	conn , err := net.ListenUDP("udp",&net.UDPAddr{IP:net.ParseIP("127.0.0.1"),Port:7878})
	if err != nil{
		panic(err)
	}
	fmt.Printf("Local <%s>\n", conn.LocalAddr().String())
	quit := make(chan struct{})
	for i := 0;i < 1;i++{
		go worker()
	}

	for i := 0;i < runtime.NumCPU();i++{
		go listener(conn,i, quit)
	}
	<-quit
}
