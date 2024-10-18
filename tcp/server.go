package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const MaxLineLenBytes = 1024
const ReadWriteTimeout = time.Minute

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Port requiered %")
	}
	host := "localhost" + ":" + os.Args[1]
	ln, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error on listening: %s\n", err)
	}
	defer ln.Close()

	log.Printf("Server running on %s\n", host)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error acepting connection from: %s\n", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	log.Printf("Accept connection: %s", conn.RemoteAddr())
	done := make(chan struct{})

	_ = conn.SetReadDeadline(time.Now().Add(ReadWriteTimeout))

	go func() {
		lim := &io.LimitedReader{
			R: conn,
			N: MaxLineLenBytes,
		}
		scan := bufio.NewScanner(lim)
		for scan.Scan() {
			in := scan.Text()
			out := strings.Fields(in)

			if _, err := conn.Write([]byte(fmt.Sprintf("Read %d commands", len(out)))); err != nil {
				log.Printf("Error on write output: %v\n", err)
				return
			}
			log.Printf("Writing: %s", out)
			lim.N = MaxLineLenBytes
			_ = conn.SetReadDeadline(time.Now().Add(ReadWriteTimeout))
		}
		done <- struct{}{}
	}()
	<-done
}
