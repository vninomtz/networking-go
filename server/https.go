package main

import (
	"fmt"
	"net/http"
)

var PORT = ":1443"

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/", Default)
	fmt.Println("Listening to port number", PORT)

	err := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("ListenAndServeTLS", err)
		return
	}
}
