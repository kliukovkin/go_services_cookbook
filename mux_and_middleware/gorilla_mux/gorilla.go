package main

import (
	"net/http"
	 "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("test", handler).Methods("GET")
	mux.
}
