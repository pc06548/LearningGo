package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	for i := 0; i < 1000; i++ {
		/*conn, err := net.Dial("tcp", ":9000")
		if err != nil {
			fmt.Println("Error creating server ", err)
		}
		fmt.Println("client started ", i)*/
		go handle(i, /*conn*/)
	}
}


func handle(i int, /*conn net.Conn*/) {

	//scanner := bufio.NewScanner(conn)
	for j := 0; j < 1000; j++ {
		//fmt.Fprintln(conn, strconv.Itoa(i) + "  "+ strconv.Itoa(j))
		fmt.Println(i, j)
		/*if scanner.Scan() {
			ln := scanner.Text()
			fmt.Println("Client the message is ", ln)
		}*/
	}
}