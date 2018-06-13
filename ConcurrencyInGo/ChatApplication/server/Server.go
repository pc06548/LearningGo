package main

import (
	"net"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var connectionCount int
var messagePool chan(string)
const (
	INPUT_BUFFER_LENGTH = 140
)

type User struct {
	Name string
	ID int
	Initiated bool
	UChannel chan []byte
	Connection *net.Conn
}

func (u *User) Listen() {
	fmt.Println("Listening for",u.Name)
	for {
		select {
		case msg := <- u.UChannel:
			fmt.Println("Sending new message to",u.Name)
			fmt.Fprintln(*u.Connection,string(msg))
		}
	}
}


type ConnectionManager struct {
	name string
	initiated bool
}

func Initiate() *ConnectionManager {
	cM := &ConnectionManager{
		name: "Chat Server 1.0",
		initiated: false,
	}
	return cM
}


func evalMessageRecipient(msg []byte, uName string) bool {
	eval := true
	expression := "@"
	re, err := regexp.MatchString(expression, string(msg))
	if err != nil {
		fmt.Println("Error:", err)
	}
	if re == true {
		eval = false
		pmExpression := "@" + uName
		pmRe, pmErr := regexp.MatchString(pmExpression, string(msg))
		if pmErr != nil {
			fmt.Println("Regex error", err)
		}
		if pmRe == true {
			eval = true
		}
	}
	return eval
}

func (cM *ConnectionManager) Listen(listener net.Listener) {
	fmt.Println(cM.name, "Started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error", err)
		}

		connectionCount++
		fmt.Println(conn.RemoteAddr(), "connected")
		user := User{Name: "anonymous", ID: 0, Initiated: false}
		Users = append(Users, &user)
		for _, u := range Users {
			fmt.Println("User online", u.Name)
		}
		fmt.Println(connectionCount, "connections active")
		go cM.messageReady(conn, &user)
	}
}

func (cM *ConnectionManager) messageReady(conn net.Conn, user *User) {
	uChan := make(chan []byte)
	for {
		buf := make([]byte, INPUT_BUFFER_LENGTH)
		n, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			conn = nil
		}
		if n == 0 {
			conn.Close()
			conn = nil
		}
		fmt.Println(n, "character message from user", user.Name)
		if user.Initiated == false {
			fmt.Println("New User is", string(buf))
			user.Initiated = true
			user.UChannel = uChan
			user.Name = string(buf[:n])
			user.Connection = &conn
			go user.Listen()
			minusYouCount := strconv.FormatInt(int64(connectionCount-1),
			10)
			conn.Write([]byte("Welcome to the chat, " + user.Name + ",there are " + minusYouCount + " other users"))
		} else {
			sendMessage := []byte(user.Name + ": " + strings.TrimRight(string(buf), " \t\r\n"))
			for _, u := range Users {
				if evalMessageRecipient(sendMessage, u.Name) == true {
					u.UChannel <- sendMessage
				}
			}
		}
	}
}

//getReady (per connectionManager) function instantiates new
//connections into a User struct, utilizing first sent message as
//the user's name.

var Users []*User
//This is our unbuffered array (or slice) of user structs.

func main() {
	connectionCount = 0
	serverClosed := make(chan bool)
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Could not start server!", err)
	}
	connManage := Initiate()
	go connManage.Listen(listener)
	<- serverClosed
}

