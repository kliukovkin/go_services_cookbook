package main

import (
	"net/http"
	"strings"
)

type httpMethod string

type routeRules struct {
	methods map[httpMethod]http.Handler
}

type router struct {
	routes map[string]routeRules
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	foundRoute, exists := r.routes[req.URL.Path]
	if exists == false {
		http.NotFound(w, req)
		return
	}
	handler, exists := foundRoute.methods[httpMethod(req.Method)]
	if exists == false {
		notAllowed(w, req, foundRoute)
		return
	}
	handler.ServeHTTP(w, req)
}

func (r *router) HandleFunc(method string, pattern string, f func(w http.ResponseWriter, req *http.Request)) {
	rules, exists := r.routes[pattern]
	if exists == false {
		rules = routeRules{methods: make(map[httpMethod]http.Handler)}
		r.routes[pattern] = rules
	}
	rules.methods[httpMethod(method)] = http.HandlerFunc(f)
}

func notAllowed(w http.ResponseWriter, req *http.Request, r routeRules) {
	methods := make([]string, 1)
	for k, _ := range r.methods {
		methods = append(methods, string(k))
	}
	w.Header().Set("Allow", strings.Join(methods, " "))
	http.Error(w, "Method Not Allowed", 405)
}

func New() *router {
	return &router{routes: make(map[string]routeRules)}
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	r := New()
	r.HandleFunc(http.MethodPost, "/test", handler)
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}
