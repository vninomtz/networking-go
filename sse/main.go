package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "SSE not supported", http.StatusInternalServerError)
			return
		}
		now := time.Now()
		log.Printf("New connection. From: %s, %s", r.Host, now.String())
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ch := make(chan string)
		go ping(r.Context(), ch)

		for val := range ch {
			event, err := encode("ping", val)
			if err != nil {
				log.Println(err)
				break
			}
			_, err = fmt.Fprint(w, event)
			if err != nil {
				log.Println(err)
				break
			}
			flusher.Flush()
		}
		log.Printf("Closing connection: %s", time.Now().String())
	})
	host := ":8080"
	log.Printf("Server listening at %s\n", host)
	http.ListenAndServe(":8080", nil)
}

func ping(ctx context.Context, ch chan<- string) {
	ticker := time.NewTicker(10 * time.Second)
outerloop:
	for {
		select {
		case <-ctx.Done():
			break outerloop
		case <-ticker.C:
			ch <- "ping"
		}
	}
	ticker.Stop()
	close(ch)
}

func encode(event string, data any) (string, error) {
	m := map[string]any{
		"data": data,
	}
	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)

	err := encoder.Encode(m)
	if err != nil {
		return "", err
	}
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", buff.String()))
	return sb.String(), nil
}
