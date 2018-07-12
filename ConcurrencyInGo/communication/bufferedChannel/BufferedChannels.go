package main

import "fmt"

var bufChan = make(chan string, 3)

func main() {

	bufChan <- "h"

	close(bufChan)

	for i := range bufChan {
		fmt.Println(i)
	}
}
