package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
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

  r := http.NewServeMux()
  r.HandleFunc("/time", timeHandler)
  r.HandleFunc("/", myHandler)

  r.HandleFunc("/debug/pprof/", pprof.Index)
  r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
  r.HandleFunc("/debug/pprof/profile", pprof.Profile)
  r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
  r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
