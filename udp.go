package main
import (
	. "fmt"
	. "net"
	"os"
	"time"

)
var channel chan *UDPConn

func NetworkListen(ch chan *UDPConn){
	udpAddr, erres := ResolveUDPAddr("udp4", "129.241.187.255:20006")
	if erres != nil {
		Println("Resolve error during listen procedure")
		os.Exit(1) 
	} 

	connection, erl := ListenUDP("udp", udpAddr)
	if erl != nil {
		Println("Listening error")
		os.Exit(1)
	}
	Println("Listening to UDP")
	ch <- connection
	
	for{
		buffer := make([]byte,1024)
		_, addr, erread := connection.ReadFromUDP(buffer)
		if erread != nil{
			Println("Reading error")
			os.Exit(1)
		}
		Println("From address: ", addr, "\nMessage: ",string(buffer))
	}
}

func NetworkSend(ch chan *UDPConn){
	sendAddr, erres := ResolveUDPAddr("udp4", "129.241.187.255:20006")
	if erres != nil {
		Println("Resolve error during send")
		os.Exit(1)
	}
	connection := <- ch

	message := []byte("Jajaja!")

	if connection == nil{
		Println("Error connection to channel")
		os.Exit(1)
	}
	for{
		connection.WriteToUDP(message, sendAddr)
		time.Sleep(1000*time.Millisecond)
	}
}



func main() {
	channel := make(chan *UDPConn, 1)
	go NetworkSend(channel)
	go NetworkListen(channel)
	for{
		time.Sleep(time.Second)
	}
}
