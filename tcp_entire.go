package main

import (
	. "fmt"
	"net"
	"os"
	"time"
)


func client(channel chan int){
	connection, err := net.Dial("tcp","129.241.187.136:33546")
	if err!= nil{
		Println("Connection error")
		os.Exit(1) 
	}
	defer connection.Close()

		p := make([]byte,1024)
		connection.Write([]byte("Connect to 129.241.187.146:20006\x00"))
		n, _ := connection.Read(p)
		Printf("We got back msg :: %s\n",p[:n])

	Println("Client done")
	channel <- 0 
}

func server(channel chan int){
	Println("Enter server")
	ln, erl := net.Listen("tcp", ":20006")
	if erl != nil{
		Println("Connection error")
		os.Exit(1) 
	}
	Println("Connection OK")
	defer ln.Close()

	connection, era := ln.Accept()
	if era != nil{
		Println("Acceptance Error")
		os.Exit(1)
	}else{
		Println("Connection Accepted")
	}
	defer connection.Close()

	p := make([]byte,1024)
	for{
		n, _ := connection.Read(p)
		Println("Connected")
			Printf("Received : %s\n",p[:n])
			connection.Write([]byte("Do you have anything to send?\x00"))
			time.Sleep(time.Second)
	}	
	channel <- 0
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go server(channel1)
	go client(channel2)

	<- channel1
	<- channel2
	Println("Program finished")
}
