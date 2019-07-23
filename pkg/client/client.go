package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func RunClient() {

	conn, _ := net.Dial("tcp", "127.0.0.1:5040")
	fmt.Println("please  import your Computational term with one operation and no white space: ")
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		fmt.Fprint(conn, input)
		if input == "end\n" {
			break
		}
		messageReceive, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("server answer:", messageReceive)
	}
}
