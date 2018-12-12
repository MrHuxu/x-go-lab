package hello

import (
	"net/rpc"

	"github.com/MrHuxu/x-go-lab/service-discovery/discovery"
)

// NewEndpoint creates an instance of discovery.Endpoint
func NewEndpoint(host string) discovery.Endpoint {
	return &endpoint{
		host: host,
	}
}

type endpoint struct {
	host string
}

func (e *endpoint) Client() (*rpc.Client, error) {
	client, err := rpc.DialHTTP("tcp", e.host)
	if err != nil {
		return nil, err
	}

	return client, nil
}
