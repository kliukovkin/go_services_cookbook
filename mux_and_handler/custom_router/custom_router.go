package main

import (
	"net/http"
	"strings"
)

type httpMethod string
type urlPattern string

type routeRules struct {
	methods map[httpMethod]http.Handler
}

type router struct {
	routes map[urlPattern]routeRules
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	foundPattern, exists := r.routes[urlPattern(req.URL.Path)]
	if !exists {
		http.NotFound(w, req)
		return
	}
	handler, exists := foundPattern.methods[httpMethod(req.Method)]
	if !exists {
		notAllowed(w, req, foundPattern)
		return
	}
	handler.ServeHTTP(w, req)
}

func (r *router) HandleFunc(method httpMethod, pattern urlPattern, f func(w http.ResponseWriter, req *http.Request)) {
	rules, exists := r.routes[pattern]
	if !exists {
		rules = routeRules{methods: make(map[httpMethod]http.Handler)}
		r.routes[pattern] = rules
	}
	rules.methods[method] = http.HandlerFunc(f)
}

func notAllowed(w http.ResponseWriter, req *http.Request, r routeRules) {
	methods := make([]string, 1)
	for k := range r.methods {
		methods = append(methods, string(k))
	}
	w.Header().Set("Allow", strings.Join(methods, " "))
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func New() *router {
	return &router{routes: make(map[urlPattern]routeRules)}
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	r := New()
	r.HandleFunc(http.MethodGet, "/test", handler)
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}
