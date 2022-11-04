package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
  fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: "
	fmt.Fprintf(w, "<h1> %s </h1>", Body)
	fmt.Fprintf(w, "<h2> %s </h2>", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	args := os.Args
  if len(args) == 1 {
    fmt.Println("Using default pot number: ", PORT)
  } else {
  	PORT = ":" + args[1]
  }

  http.HandleFunc("/time", timeHandler)
  http.HandleFunc("/", myHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
