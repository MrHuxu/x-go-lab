package engine

import (
	"log"
	"net/http"
	"strconv"
)

type Engine struct {
	router *Router
}

func (e *Engine) Run(port int) {
	str := strconv.Itoa(port)
	log.Fatal(http.ListenAndServe(":"+str, e))
}

func DefaultEngine() *Engine {
	engine := &Engine{
		router: &Router{
			handlers: make(map[string]Handler),
			children: map[string]*Router{"/": defaultIndexRouter},
		},
	}
	return engine
}
