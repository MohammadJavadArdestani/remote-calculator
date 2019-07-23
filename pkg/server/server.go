package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func RunServer() {

	fmt.Println("server is running")
	client, err := net.Listen("tcp", ":5040")
	if err != nil {
		fmt.Println("1")
	}
	for {
		connection, err := client.Accept()
		if err != nil {
			fmt.Println("2")
		}
		go clientHandler(connection)

	}
}
func clientHandler(connection net.Conn) {
	for {
		receivedMessage, _ := bufio.NewReader(connection).ReadString('\n')
		//if you want run it on linux use "\n"
		receivedMessage = strings.TrimRight(receivedMessage, "\r\n")
		if receivedMessage == "end" {
			connection.Close()
			fmt.Println("client disconnected")
			fmt.Println("server is waiting for new client")
			break
		}
		result := Calculate(receivedMessage)

		answer := strconv.FormatFloat(result, 'f', -1, 64)

		_, _ = fmt.Fprint(connection, answer+"\r\n")

	}
}
