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
	for {
		connection, err := client.Accept()
		if err != nil {
			fmt.Println("2")
		}
		for {
			recivedMessage, _ := bufio.NewReader(connection).ReadString('\n')
			if recivedMessage[len(recivedMessage)-1] == '\n' {
				recivedMessage = recivedMessage[:len(recivedMessage)-1]
			}
			if recivedMessage == "end" {
				connection.Close()
				fmt.Println("client disconnected")
				break

			}
			result := Calculate(recivedMessage)

			answer := strconv.FormatFloat(result, 'f', -1, 64)

			_, _ = fmt.Fprint(connection, answer+"\n")

		}
	}
}
