package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	url := os.Args[1]

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}

	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	s := strings.TrimSpace(string(content))
	fmt.Println(s)

}
