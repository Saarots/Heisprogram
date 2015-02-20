package main

import (
	."fmt"
	"time"
	."net"
	"os"
	"encoding/json"
)

var i = 0
var masterMode = 0

type Data struct{
	Message string
	Number int
}
func master() {
	sendAddr, err := ResolveUDPAddr("udp", ":31004") // bordcaster
	if err != nil{
		Println("Resolve error during send")
		os.Exit(1)
	}
	con, _ := DialUDP("udp", nil, sendAddr) // lager socket til sending
	defer con.Close()// stenger porten etter funksjonene
	i++
	Println("Jeg er i MasterMode!\ni = ",i)
	dataSend := Data{
		Message: "Alive",
		Number: i,
	}
	messageSend, _ := json.Marshal(dataSend)
	Fprintf(con, string(messageSend))// Triks for å sende til UDP(tar KUN string)
	Println("\nMelding sendt")
	time.Sleep(1*time.Second)
}
func slave() {
	Println("I am in SlaveMode!\ni:  %v", i)
	udpAddr, err := ResolveUDPAddr("udp", ":31004")
	if err != nil{
		Println("Resolve error during listen procedure")
		os.Exit(1)
	}
	ln, err := ListenUDP("udp", udpAddr) //lager socket for høring
	if err != nil{
		Println("Listening error")
		os.Exit(1)
	}
	for{
		ln.SetReadDeadline(time.Now().Add(3*time.Second))
		buffer := make([]byte, 1024)
		n, err := ln.Read(buffer) // test "_" på de andre etter på for de vi ikke trenger
		if err != nil{
			masterMode = 1
			Println("I am Master now!!!!")
			ln.Close()
			break
		}
		var messageRecv []Data
		json.Unmarshal(buffer[0:n], &messageRecv) // lager i ny struct så er det bare å hente fra den
		//msg := string(messageRecv)
		Println("Message recived is: %v", messageRecv)
		// printer ikke rett så må få endra det!!!!
		//i = messageRecv
	}
}

func main() {
	masterMode = 0
	for{
		if masterMode == 1{
			master()
		}else{
			slave()
		}
	}
}
