package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	connect := os.Args[1]
	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port")
	}
	conn, err := net.Dial("tcp", connect)
	if err != nil {
		log.Printf("Error %s\n", err)
		return
	}
	defer conn.Close()
	//go onMessage(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
		onMessage(conn)

	}
}

func onMessage(conn net.Conn) {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("Conn closed %s\n", err)
		return
	}
	fmt.Printf("--> %s", msg)
	fmt.Print(">")
}
