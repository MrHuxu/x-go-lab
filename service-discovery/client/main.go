package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MrHuxu/x-go-lab/service-discovery/discovery"
	"github.com/MrHuxu/x-go-lab/service-discovery/proto/hello"
)

var (
	service = "Server.SayHello"
)

func main() {
	d := discovery.New(hello.NewEndpoint)

	for {
		fmt.Println("\x1Bc")

		for i := 0; i < 10; i++ {
			endpoint := d.Get(strconv.Itoa(i))
			client := endpoint.Client()

			var (
				args  = hello.Args{Seq: i}
				reply hello.Reply
			)
			if err := client.Call(service, args, &reply); err != nil {
				log.Fatal("arith error:", err)
			}
			fmt.Println(args, reply)
		}

		time.Sleep(time.Second * 2)
	}
}
