package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", handler).Methods("GET")
	http.ListenAndServe(":8000", r)
}
