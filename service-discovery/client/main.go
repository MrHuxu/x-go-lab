package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/MrHuxu/x-go-lab/service-discovery/discovery"
	"github.com/MrHuxu/x-go-lab/service-discovery/proto/hello"
)

var (
	client  *rpc.Client
	service = "Server.SayHello"
)

func main() {
	getClient()
	invokeService()
}

func getClient() {
	d := discovery.New(hello.NewEndpoint)
	endpoint := d.Get("1")
	client = endpoint.Client()
}

func invokeService() {
	var (
		args  = hello.Args{Seq: 1}
		reply hello.Reply
	)
	if err := client.Call(service, args, &reply); err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Println(args, reply)
}
