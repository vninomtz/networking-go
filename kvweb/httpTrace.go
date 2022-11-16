package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main()  {
  if len(os.Args) != 2 {
    fmt.Printf("Usage: URL\n")
    return
  }
  URL := os.Args[1]
  client := http.Client{}
  req, _ := http.NewRequest("GET", URL, nil)
  trace := &httptrace.ClientTrace{
    GotFirstResponseByte: func() {
      fmt.Println("First response byte!")
    },
    GotConn: func(gci httptrace.GotConnInfo) {
      fmt.Printf("Got Conn: %+v\n", gci)
    },
    DNSDone: func(di httptrace.DNSDoneInfo) {
      fmt.Printf("DNS info: %+v\n", di)
    },
    ConnectStart: func(network, addr string) {
      fmt.Println("Dial start")
    },
    ConnectDone: func(network, addr string, err error) {
      fmt.Println("Dial done")
    },
    WroteHeaders: func() {
      fmt.Println("Wrote headers")
    },
  }
  req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
  fmt.Println("Requesting data from server!")
  _, err := http.DefaultTransport.RoundTrip(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  resp, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  io.Copy(os.Stdout, resp.Body)
}
