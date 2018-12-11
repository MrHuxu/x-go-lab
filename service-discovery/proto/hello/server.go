package hello

import (
	"fmt"
)

type Args struct {
	Seq int
}

type Reply struct {
	Msg string
}

type Server struct {
	Sign string
}

func (s *Server) SayHello(args *Args, rep *Reply) error {
	rep.Msg = fmt.Sprintf("This is a hello from server %s to client %d", s.Sign, args.Seq)
	return nil
}
