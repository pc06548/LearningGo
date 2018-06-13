package main

import (
	"net"
	"fmt"
	"bufio"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Error creating server ", err)
	}

	fmt.Println("server started")
	for {
		conn, err := listener.Accept()
		fmt.Println("new connection received")
		if err != nil {
			fmt.Println("Error in listening ", err)
		}
		go handle(conn)
	}



}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	fmt.Println("handling connection")
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		fmt.Fprintln(conn, "Copied you ",ln, " tell me more..")
	}
}
