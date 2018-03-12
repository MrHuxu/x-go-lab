package engine

import (
	"net/http"
	"strings"
)

const routeUndefined = "Route undefined"

type Handler func(http.ResponseWriter, *http.Request)

type Router struct {
	handlers map[string]Handler
	children map[string]*Router
}

var defaultIndexRouter = &Router{
	handlers: map[string]Handler{
		"GET": func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(routeUndefined))
		},
	},
	children: make(map[string]*Router),
}

func (e *Engine) Get(path string, handler Handler) {
	e.storeRoute(getPatterns(path), "GET", handler)
}

func (e *Engine) Post(path string, handler Handler) {
	e.storeRoute(getPatterns(path), "POST", handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	patterns := getPatterns(r.URL.Path)
	router := e.router.children["/"]
	for _, pattern := range patterns {
		if _, ok := router.children[pattern]; !ok {
			w.Write([]byte(routeUndefined))
			return
		}
		router = router.children[pattern]
	}

	if _, ok := router.handlers[r.Method]; !ok {
		w.Write([]byte(routeUndefined))
		return
	}
	router.handlers[r.Method](w, r)
}

func getPatterns(path string) []string {
	patternString := strings.Trim(path, "/")
	if len(patternString) == 0 {
		return []string{}
	}
	return strings.Split(patternString, "/")
}

func (e *Engine) storeRoute(patterns []string, method string, handler Handler) {
	var router = e.router.children["/"]
	for _, pattern := range patterns {
		_, ok := router.children[pattern]
		if !ok {
			router.children[pattern] = &Router{
				handlers: make(map[string]Handler),
				children: make(map[string]*Router),
			}
		}
		router = router.children[pattern]
	}
	router.handlers[method] = handler
}
