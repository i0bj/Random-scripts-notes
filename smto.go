
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	var mailServer string
	fmt.Println("Hello, please enter the mail server IP")
	fmt.Scanln(&mailServer)
	var email string
	fmt.Println("Next, enter the email address to verify: ")
	fmt.Scanln(&email)

	// Create socket/connection to mail server
	conn, err := net.Dial("tcp", mailServer+":25")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// Sending VRFY cmd across the wire
	conn.Write([]byte("VRFY " + email + `\r\n`))
	log.Printf("Sending: VRFY cmd")

	results, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(results))
}
