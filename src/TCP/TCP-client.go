package main

import (
	"fmt"
	"bufio"
	"os"
	"net"
)

func main() {

	// Set up connection for

	// reader := bufio.NewReader(os.Stdin).ReadString('\n')
	// fmt.Print("Enter host:port")
	conn, err := net.Dial("tcp", "127.0.0.1:12000")
	if err != nil {
    	fmt.Println("Error reading:", err.Error())
  	}

	for {
		// read user input
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
    		fmt.Println("Error reading:", err.Error())
 		}

		// send input to socket
		fmt.Fprintf(conn, text + "\n")

		// listen
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// print message
		fmt.Print("Message from server: " + message)
	}



}