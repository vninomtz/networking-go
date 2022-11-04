package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
  args := os.Args
  if len(args) == 1 {
    fmt.Println("Need a domain name!")
    return
  }
  domain := args[1]
  MXs, err := net.LookupMX(domain)
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, MX := range MXs {
    fmt.Println(MX.Host)
  }
}
