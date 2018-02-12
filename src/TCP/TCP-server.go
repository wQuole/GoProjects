package main

import (
	"fmt"
	"strings" // for processing message
	"net"
	"bufio"
	"os"
)

func main() {

	fmt.Println("Server launching...")


	// Listener for incoming connections
	ln, err := net.Listen("tcp",  ":12000")
	if err != nil {
		fmt.Println("Error occured:", err.Error())
		os.Exit(1)
	}

	// accept connetion 
	conn, err := ln.Accept()

	// loop
	for {

		// listen for message
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// print message
		fmt.Print("Message recieved...\n", string(message))

		// process message 
		newMessage := strings.ToUpper(message)

		// send processed message
		conn.Write([]byte(newMessage + "\n"))

	}
}


