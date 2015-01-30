package main

import (
	. "fmt"
	"net"
	"os"
)

func main(){
	connection, err := net.Dial("tcp","129.241.187.136:33546")
	if err!= nil{
		Println("Connection error")
		os.Exit(1) 
	}
	defer connection.Close()

	p := make([]byte,2048)
	connection.Write([]byte("Connect to 129.241.187.146:20006"))
	n, _ := connection.Read(p)
	Printf("We got back msg :: %s\n",p[:n])
	Println("Done")
}
