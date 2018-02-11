package web

import (
	"net/http"
	"strings"
)

type Handler func(http.ResponseWriter, *http.Request)

type EndPoint struct {
	Funcs    map[string]Handler
	Children map[string]*EndPoint
}

var defaultIndexEntPoint = &EndPoint{
	Funcs: map[string]Handler{
		"GET": func(w http.ResponseWriter, r *http.Request) {
			panic("Router undefined")
		},
	},
}

func (e *Engine) Get(path string, handler Handler) {
	patternString := strings.Trim(path, "/")
	if len(patternString) == 0 {
		e.storeRoute([]string{}, "GET", handler)
	} else {
		patterns := strings.Split(patternString, "/")
		e.storeRoute(patterns, "GET", handler)
	}
}

func (e *Engine) Post(path string, handler Handler) {
	patternString := strings.Trim(path, "/")
	if len(patternString) == 0 {
		e.storeRoute([]string{}, "POST", handler)
	} else {
		patterns := strings.Split(patternString, "/")
		e.storeRoute(patterns, "POST", handler)
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println("test test")
}

func (e *Engine) storeRoute(patterns []string, method string, handler Handler) {
	var endPoint = e.Router.Children["/"]
	for _, pattern := range patterns {
		_, ok := endPoint.Children[pattern]
		if !ok {
			endPoint.Children[pattern] = &EndPoint{
				Funcs:    make(map[string]Handler),
				Children: make(map[string]*EndPoint),
			}
		}
		endPoint = endPoint.Children[pattern]
	}
	endPoint.Funcs[method] = handler
}
