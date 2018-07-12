package main

import "fmt"

func main() {
	for i := 0; i< 100000; i++ {
		go loop()
	}
}

func loop() {
	fmt.Println("Hi there")
}
