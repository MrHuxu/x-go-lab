package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"

	"github.com/MrHuxu/x-go-lab/service-discovery/discovery"
	"github.com/MrHuxu/x-go-lab/service-discovery/proto/hello"
)

var (
	port string
	sign string
)

func main() {
	parseArgs()
	launchServer()
	hangServer()
}

func parseArgs() {
	flag.StringVar(&port, "port", "", "port of the server")
	flag.StringVar(&sign, "sign", "", "sign of the server")
	flag.Parse()
}

func launchServer() {
	helloServer := &hello.Server{Sign: sign}
	rpc.Register(helloServer)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":"+port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	discovery.Register("localhost:" + port)
}

func hangServer() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-exit:
		println("server stopped.")
		discovery.Unregister("localhost:" + port)
		println("service unregistered.")
	}
}
