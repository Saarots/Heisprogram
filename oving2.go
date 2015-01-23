package main
import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
			// This is probably not a good idea for large projects...
	"runtime"
	//"time"
)
var testverdi int = 0;
var channel = make(chan int, 1)
var add_done = make(chan int, 1)
var subtract_done = make(chan int, 1) 


func adder(){
	for j:=0; j <999999; j++{ //999999 1000000
 		<- channel
		testverdi++
		channel <- 1
	}
	add_done <- 1
}

func subtractor(){
	for k:=0; k <1000000; k++{
		<- channel  
		testverdi--
		channel <- 1
	}
	subtract_done <- 1
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	channel <- 1

	go adder()
	go subtractor()
	<- add_done
	<- subtract_done
	//time.Sleep(100*time.Millisecond)
	Println("Testvariabelen er: ",testverdi)
}
