// Package mux implements a simple HTTP request multiplexer.
package mux

import (
	"net/http"
	"strings"
)

// Router matches incoming requests to registered handlers.
type Router struct {
	routes []route
}

type route struct {
	pattern string
	handler http.Handler
}

// New returns an empty Router.
func New() *Router {
	return &Router{}
}

// Handle registers handler for the given pattern.
func (r *Router) Handle(pattern string, handler http.Handler) {
	r.routes = append(r.routes, route{pattern: pattern, handler: handler})
}

// HandleFunc registers a function as a handler for the given pattern.
func (r *Router) HandleFunc(pattern string, fn http.HandlerFunc) {
	r.Handle(pattern, fn)
}

// ServeHTTP dispatches the request to the first matching route.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, rt := range r.routes {
		if strings.HasPrefix(req.URL.Path, rt.pattern) {
			rt.handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
