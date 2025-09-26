package main

import (
	"flag"
	"io"
	"log"
	"net"
)

func main() {
	port := flag.String("port", ":7", "Listen port")
	flag.Parse()

	listener, err := net.Listen("tcp", *port)
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
