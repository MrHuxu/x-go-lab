package discovery

import (
	"net/rpc"
)

type Endpoint interface {
	Client() *rpc.Client
}
