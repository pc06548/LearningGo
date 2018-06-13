package main

import (
	"time"
	"sync"
	"fmt"
	"runtime"
)

func wait() {
	time.Sleep(5*time.Second)
	wg.Done()
}
var wg = sync.WaitGroup{}

func main()  {
	println(runtime.GOMAXPROCS(1))
	startTime := time.Now().Second()
	wg.Add(1000000)
	for i := 0; i<1000000; i++ {
		fmt.Println(i)
		go wait()
	}
	wg.Wait()
	endTime := time.Now().Second()
	fmt.Println("done last")
	fmt.Println(startTime - endTime)
}
