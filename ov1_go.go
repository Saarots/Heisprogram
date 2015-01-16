package main
import (
	. "fmt" // Using '.' to avoid prefixing functions with their package names
		// This is probably not a good idea for large projects...
	"runtime"
	"time"
)
var i = 0;
func Thread1func(){
	for j:=0; j <1000000; j++{
		i++;
	}
}


func Thread2func(){
	for k:=0; k <1000000; k++{
		i--;
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	go Thread1func()
	//time.Sleep(100*time.Millisecond)
	go Thread2func()
	time.Sleep(100*time.Millisecond)

	Println(i)
}
