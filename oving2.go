package main
import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
			// This is probably not a good idea for large projects...
	"runtime"
	//"time"
)
var TestVerdi = 0;

func Thread1func(channel chan int){
	for j:=0; j <1000000; j++{
		i := <- channel
		i++;
		channel <- i
	}
}

func Thread2func(channel chan int){
	for k:=0; k <1000000; k++{
		i := <- channel  
		i--;
		channel <- i
	}
}

func main() {
	
	channel := make(chan int, 1)
	channel <- TestVerdi
	runtime.GOMAXPROCS(runtime.NumCPU())

	go Thread1func(channel)
	go Thread2func(channel)
	TestVerdi := <- channel
	//time.Sleep(100*time.Millisecond)
	Println(TestVerdi)
}
