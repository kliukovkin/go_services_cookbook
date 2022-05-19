package main

import "net/http"

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
	}
	w.Write([]byte("New user has been created"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", createUser)
	http.ListenAndServe(":3000", mux)
}
