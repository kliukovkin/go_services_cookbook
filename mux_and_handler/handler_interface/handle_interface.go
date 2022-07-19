package main

import (
	"fmt"
	"net/http"
)

type handler struct{}

func (t *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping %s\n", r.URL.Query().Get("name"))
}

/*
❯ curl "localhost:8000?name=foo"
ping foo
*/
func main() {
	h := &handler{}
	http.Handle("/", h)
	http.ListenAndServe(":8000", nil)

}
