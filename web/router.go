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

var defaultIndexEndPoint = &EndPoint{
	Funcs: map[string]Handler{
		"GET": func(w http.ResponseWriter, r *http.Request) {
			panic("Router undefined")
		},
	},
	Children: make(map[string]*EndPoint),
}

func (e *Engine) Get(path string, handler Handler) {
	e.storeRoute(getPatterns(path), "GET", handler)
}

func (e *Engine) Post(path string, handler Handler) {
	e.storeRoute(getPatterns(path), "POST", handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	patterns := getPatterns(r.URL.Path)
	endPoint := e.Router.Children["/"]
	for _, pattern := range patterns {
		if _, ok := endPoint.Children[pattern]; !ok {
			w.Write([]byte("Router not found"))
			return
		}
		endPoint = endPoint.Children[pattern]
	}

	if _, ok := endPoint.Funcs[r.Method]; !ok {
		w.Write([]byte("Router not found"))
		return
	}
	endPoint.Funcs[r.Method](w, r)
}

func getPatterns(path string) []string {
	patternString := strings.Trim(path, "/")
	if len(patternString) == 0 {
		return []string{}
	}
	return strings.Split(patternString, "/")
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
