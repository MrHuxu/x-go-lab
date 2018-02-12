package main

import (
	"net/http"

	"github.com/MrHuxu/x-go-lab/web"
)

func main() {
	svr := web.DefaultEngine()

	svr.Get("/test/get", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test get work"))
	})
	svr.Post("/test/post", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test post work"))
	})

	svr.Run(11101)
}
