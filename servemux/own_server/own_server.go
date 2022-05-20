package own_server

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping %s\n", r.URL.Query().Get("name"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
		// ReadTimeout is the maximum duration for reading the entire
		// request, including the body.
		ReadTimeout: 5 * time.Second,
		// WriteTimeout is the maximum duration before timing out
		// writes of the response.
		WriteTimeout: 10 * time.Second,
	}
	srv.ListenAndServe()
}
