package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/MrHuxu/x-go-lab/service-discovery/discovery"
	"github.com/MrHuxu/x-go-lab/service-discovery/proto/hello"
)

var (
	funcName = "Service.SayHello"
)

func main() {
	d := discovery.New(hello.NewEndpoint)

	for {
		fmt.Println("\x1Bc")

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func(i int) {
				defer wg.Done()
				invokeService(d, i)
			}(i)
		}

		wg.Wait()
		time.Sleep(time.Second * 2)
	}
}

func invokeService(d discovery.Discovery, seq int) {
	endpoint, err := d.Get(strconv.Itoa(seq))
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := endpoint.Client()
	if err != nil {
		fmt.Println(err)
		return
	}

	var (
		args  = hello.Args{Seq: seq}
		reply hello.Reply
	)
	if err := client.Call(funcName, args, &reply); err != nil {
		fmt.Println(err)
	}
	fmt.Println(args, reply)
}
