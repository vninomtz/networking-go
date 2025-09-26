package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// https://www.rfc-editor.org/rfc/rfc862.html
// TODO: Implement echo using UDP

func main() {
	port := flag.String("port", ":7", "Listen port")
	serve := flag.Bool("serve", false, "Run the ECO Server")
	flag.Parse()

	if *serve {
		Run(*port)
		return
	}
	msg := strings.Join(os.Args[1:], " ")
	Echo(*port, msg)

}

func Run(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error creating listener: %w", err)
	}
	defer listener.Close()

	log.Printf("TCP Server listening at %s\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error to accept connection: %w\n", err)
			continue
		}
		go handleWithCopy(conn)
	}
}

func Echo(port, msg string) {
	conn, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatalf("Error create dial: %w", err)
	}
	defer conn.Close()

	start := time.Now()
	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Fatalf("Error to write: %w", err)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatalf("Error to read: %w", err)
	}
	rtt := time.Since(start)
	fmt.Printf("ECHO: %s\n", string(buf))
	fmt.Printf("RTT: %v", rtt)
}

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1028)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %w\n", err)
			}
			log.Printf("Closing connection from %s\n", conn.RemoteAddr())
			return
		}

		log.Printf("Received from %s: %s\n", conn.RemoteAddr(), string(buf))

		conn.Write(buf)
	}
}
func handleWithCopy(conn net.Conn) {
	defer conn.Close()

	log.Printf("New connection: %s\n", conn.RemoteAddr())
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Printf("Error copying from connection: %w\n", err)
	}
	log.Printf("Close connection: %s\n", conn.RemoteAddr())
}
