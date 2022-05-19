package servemux

import (
	"fmt"
	"net/http"
)

type handler struct{}

/*
We need to implement interface:
type Handler interface{
   ServeHTTP(ResponseWriter, *Request)
}
to achieve that we need to provide our handler struct with single method
ServeHTTP(w http.ResponseWriter, r *http.Request)
*/

func (t *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping %s\n", r.URL.Query().Get("name"))
}
func main() {
	h := &handler{}
	http.Handle("/", h)
	http.ListenAndServe(":8000", nil)

}
