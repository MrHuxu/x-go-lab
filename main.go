package main

import (
	"github.com/MrHuxu/x-go-lib/web"
)

func main() {
	svr := web.DefaultServer()
	svr.Run(11101)
}
