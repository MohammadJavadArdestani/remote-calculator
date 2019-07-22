package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

func RunServer() {
 fmt.Println("server is running")
	client, err := net.Listen("tcp", ":5040")
	if err !=nil{
		fmt.Println("1")
	}
	connection, err := client.Accept()
	if err !=nil {
		fmt.Println("2")
	}
	recivedMessage, _ := bufio.NewReader(connection).ReadString('\n')
	result:= Calculate(recivedMessage)
	//fmt.Println(result)
	//fmt.Println("its showing on server:",result)
	//result := strings.ToLower(recivedMessage)
	answer:=strconv.FormatFloat(result,'f',-1,64)

	_, _ = fmt.Fprint(connection, answer+"\n")
}
	/*fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))*/



