package web

import (
	"log"
	"net/http"
	"strconv"
)

type Engine struct {
	Port   int
	Router *EndPoint
}

func (e *Engine) Run(port int) {
	e.Port = port

	str := strconv.Itoa(port)
	log.Fatal(http.ListenAndServe(":"+str, e))
}

func DefaultServer() *Engine {
	engine := &Engine{
		Router: &EndPoint{
			Funcs:    make(map[string]Handler),
			Children: map[string]*EndPoint{"/": defaultIndexEntPoint},
		},
	}
	return engine
}
