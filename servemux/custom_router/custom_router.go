package custom_router

import (
	"fmt"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/foo":
		fmt.Fprint(w, "here is /foo")
	case "/bar":
		fmt.Fprint(w, "here is /bar")
	case "/baz":
		fmt.Fprint(w, "here is /baz")
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}
