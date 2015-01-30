package main

import (
	. "fmt"
	"net"
	"os"
)

func handleConnections(connection net.Conn) {
	Println("Connected!")
	p := make([]byte,2048)
	n, _ := connection.Read(p)
	Printf("Received : %s\n",p[:n])
}

func main() {
	Println("Start")
	ln, err := net.Listen("tcp", ":20006")
	if err != nil{
		Println("Connection error")
		os.Exit(1) 
	}
	defer ln.Close()
	for{
		connection, err := ln.Accept()
		Println("Accepted")
		if err != nil{
			Println("Error")
			continue
		}
		go handleConnections(connection)
	}
	Println("Done")
}
