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

var timeout = make(chan int, 1) 
var number = make(chan int, 1)

type Data struct{
	Message string
	Number int
}



func sendNumber() {
	sendAddr, err := ResolveUDPAddr("udp", ":31004") // bordcaster
	if err != nil{
		Println("Resolve error during send")
		os.Exit(1)
	}
	con, _ := DialUDP("udp", nil, sendAddr) // lager socket til sending
	defer con.Close()// stenger porten etter funksjonene

	<- number
	dataSend := Data{
		Message: "Alive",
		Number: i,
	}
	
	messageSend, _ := json.Marshal(dataSend)
	os.Stdout.Write(messageSend)	
	number <- 1
	Fprintf(con, string(messageSend))// Triks for å sende til UDP(tar KUN string)
		Println("\nMelding sendt")


}

func slave() {
	Println("I am in SlaveMode")
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

	Println("Hører etter beskjed")
	
	for{
		Println("Starter å høre")
		ln.SetReadDeadline(time.Now().Add(3*time.Second))
		buffer := make([]byte, 1024)
		_, addr, err := ln.ReadFromUDP(buffer) // test "_" på de andre etter på for de vi ikke trenger
		if err != nil{
			masterMode = 1
			Println("I am Master:", masterMode)
			ln.Close()
			break
		}
		Println("Printer ut beskjed")
		var messageRecv []Data
		json.Unmarshal(buffer[], &messageRecv)
		Println("From address: ", addr, "\nMessage: ", messageRecv)
		// printer ikke rett så må få endra det!!!!
		//i = messageRecv
		
	}
}

func master(){
	<- number
	i++
	number <- i
	sendNumber()
	Println("Jeg er i MasterMode!\ni = ",i)
	time.Sleep(1*time.Second)
}


func main() {
	masterMode = 0
	number <- 1

	for{
		if masterMode == 1{
			master()
		}else{
			slave()
		}
	}
}
