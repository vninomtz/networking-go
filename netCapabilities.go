package main

// Purpose: revel the capabilities of each network
// interface found on your UNIX system
import (
	"fmt"
	"net"
)

// net.Interface structure

// type Interface struct {
//   Index int
//   MTU int
//   Name string
//   HardwareAddr HardwareAddr
//   Flags Flags
// }

func main()  {
  interfaces, err := net.Interfaces()

  if err != nil {
    fmt.Print(err)
    return
  }

  for _, i := range interfaces {
    fmt.Printf("Name: %v\n", i.Name)
    fmt.Println("Interface Flags: ", i.Flags.String())
    fmt.Println("Interface MTU: ", i.MTU)
    fmt.Println("Interface Hardware Address: ", i.HardwareAddr)

    fmt.Println()
  }
}
