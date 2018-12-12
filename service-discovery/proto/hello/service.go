package hello

import (
	"fmt"
)

// Args defines the fields of the args
type Args struct {
	Seq int
}

// Reply defines the fields of the reply
type Reply struct {
	Msg string
}

// Service defines the fields of the service
type Service struct {
	Sign string
}

// SayHello ...
func (s *Service) SayHello(args *Args, rep *Reply) error {
	rep.Msg = fmt.Sprintf("This is a message from server %s to client %d", s.Sign, args.Seq)
	return nil
}
