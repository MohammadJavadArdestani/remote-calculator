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
	if err != nil {
		fmt.Println("1")
	}
	connection, err := client.Accept()
	if err != nil {
		fmt.Println("2")
	}
	recivedMessage, _ := bufio.NewReader(connection).ReadString('\n')
	result := Calculate(recivedMessage)

	answer := strconv.FormatFloat(result, 'f', -1, 64)

	_, _ = fmt.Fprint(connection, answer+"\n")
}
