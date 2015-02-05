package main

import(
	"fmt"
	"net"
	"os"
	"time"
	"runtime"
)

const MY_IP = "129.241.1887.146"
const MY_PORT = "20006"
const TARGET_PORT = "33546"
const TARGET_IP = "129.241.187.136"

func connectToServer(ipAddr string, port string, channel chan bool){
	serverAddr, err := net.ResolveTCPAddr("tcp",ipAddr+":"+port)
	checkError(err)

	con, err := net.DialTCP("tcp", nil, serverAddr);
	checkError(err)

	cd := make([]byte,1024)
	con.Read(cd)
	fmt.Printf("%s",cd)
	stop := 0

	msg := "Connecting to: "+MY_IP+":"+MY_PORT+"\x00"
	con.Write([]byte(msg))

	for stop != -1{
		input := ""
		fmt.Scanf("%s",&input)
		if input=="stop"{
			stop = -1
		}
		_,err := con.Write([]byte(input+"\x00"))
		checkError(err)

		_,err = con.Read(cd)
		checkError(err)

		fmt.Printf("%s\n",cd)
	}
	channel <- true
}

func ListenConnection(port string, channel chan bool){
	psock, err := net.Listen("tcp", ":"+port)
	checkError(err)

	conn, err := psock.Accept()
	checkError(err)

	go func(conn net.Conn){
		for{
			buffer := make([]byte,1024)
			_,err := conn.Read(buffer)
			checkError(err)
			fmt.Printf("%s\n",buffer)
		}
	}(conn)
	for{
		msg := "Now it should work\x00"
		_,err := conn.Write([]byte(msg))
		checkError(err)
		time.Sleep(2*time.Second)
	}
	channel <- true
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	channel := make(chan bool, 1)

	go connectToServer(TARGET_IP, TARGET_PORT, channel)
	go ListenConnection(MY_PORT, channel)

	
	<- channel
}
