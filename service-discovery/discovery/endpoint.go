package discovery

import (
	"net/rpc"
)

// Endpoint defines the interface that a service endpoint needs to implement
type Endpoint interface {
	Client() (*rpc.Client, error)
}
