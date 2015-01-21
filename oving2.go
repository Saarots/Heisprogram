package main
import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
	// This is probably not a good idea for large projects...
	"runtime"
	"time"
)
var i = 0;
var cs = make(chan int, 1)
var addDone = make(chan string)
func Thread1func(){

	for j:=0; j <1000000; j++{
		cs <- 1		
		i++;
		cs <- 0
	}
	addDone <- "Done incrementing"
}
func Thread2func(){

	for k:=0; k <1000000; k++{
		cs <- 1
		i--;
		cs <- 0
	}
	addDone <- "Done decrementing"
}

func main() {
	cs <- 0
	runtime.GOMAXPROCS(runtime.NumCPU())

	go Thread1func()
	go Thread2func()
	time.Sleep(100*time.Millisecond)
	println(addDone)	
	println(addDone)
	Println(i)
}
