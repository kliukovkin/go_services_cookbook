package handle_func

import (
	"fmt"
	"net/http"
)

/*
Same function signature as for HandlerServeHTTP, just a syntax sugar
*/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping %s\n", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
