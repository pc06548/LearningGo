package main


import (
	"fmt"
	"runtime" )
func listThreads() int {
//returns previous setting.
	threads := runtime.GOMAXPROCS(0)
	return threads
}
func main() {
	//runtime.GOMAXPROCS(4)
	fmt.Printf("%d thread(s) available to Go.", listThreads())
}