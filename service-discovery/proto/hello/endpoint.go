package hello

import (
	"log"
	"net/rpc"

	"github.com/MrHuxu/x-go-lab/service-discovery/discovery"
)

func NewEndpoint(host string) discovery.Endpoint {
	return &endpoint{
		host: host,
	}
}

type endpoint struct {
	host string
}

func (e *endpoint) Client() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", e.host)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}
