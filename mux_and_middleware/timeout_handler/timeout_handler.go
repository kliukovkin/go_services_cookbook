package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "ping %s\n", r.URL.Query().Get("name"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.TimeoutHandler(http.HandlerFunc(handler), time.Second * 1, "Timeout"))
	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
		WriteTimeout: 2 * time.Second,
	}
	srv.ListenAndServe()
}
