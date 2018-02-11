package main

import (
	"github.com/MrHuxu/x-go-lab/web"
)

func main() {
	svr := web.DefaultServer()
	svr.Run(11101)
}
