package main

import (
	"fmt"
	"net"
	"os"
)

// If the given command-line argument is a valid IP address,
// the program will process it as an IP address; otherwise
// it will assume that it is dealing with a hostname that needs
// to be translated into one or more IP addresses.

func lookIP(address string) ([]string, error) {
  hosts, err := net.LookupAddr(address)
  if err != nil {
    return nil, err
  }
  return hosts, nil
}

func lookHostname(hostname string) ([]string, error) {
  IPs, err := net.LookupHost(hostname)
  if err != nil {
    return nil, err
  }
  return IPs, nil
}

func main()  {
  args := os.Args
  if len(args) == 1 {
    fmt.Println("Please provide an argument!")
    return
  }
  input := args[1]
  IPaddrs := net.ParseIP(input)

  if IPaddrs == nil {
    IPs, err := lookHostname(input)
    if err == nil {
      for _, singleIp := range IPs {
        fmt.Println(singleIp)
      }
    }
  } else {
    hosts, err := lookIP(input)
    if err == nil {
      for _, hostname := range hosts {
        fmt.Println(hostname)
      }
    }
  }
}
