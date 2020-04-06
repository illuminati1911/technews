package utils

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPC is a utility for creating GRPC server.
type GRPC struct {
	Server   *grpc.Server
	listener net.Listener
}

// NewGRPC creates an instance of GRPC with given port
func NewGRPC(port string) *GRPC {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	return &GRPC{server, listener}
}

// Start will start the GRPC server at given port
func (g *GRPC) Start() error {
	reflection.Register(g.Server)
	return g.Server.Serve(g.listener)
}
