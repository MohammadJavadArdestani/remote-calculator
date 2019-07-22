package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}

	return strings.FieldsFunc(s, splitter)
}*/

func RunClient() {
	/*reader:=bufio.NewReader(os.Stdin)
	clientInput , _:=reader.ReadString ('\n')

	s := SplitAny(clientInput,"+-")

	ip, port := s[0], s[1]

	fmt.Print(ip+"annnnd"+port)
	fmt.Print("saaalamm")*/


		// connect to this socket
		conn,_ :=net.Dial("tcp","127.0.0.1:5040")
		fmt.Println("please your import Computational term whit one operation and no white space: ")
		reader:=bufio.NewReader(os.Stdin)
		input,_:=reader.ReadString('\n')
		fmt.Fprint(conn,input)
		messageRecive,_:=bufio.NewReader(conn).ReadString('\n')
		fmt.Println("server answer:",messageRecive)

		/*conn, _ := net.Dial("tcp", "127.0.0.1:8081")
		for {
			// read in input from stdin
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			// send to socket
			fmt.Fprintf(conn, text + "\n")
			// listen for reply
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: "+message)*/

}
