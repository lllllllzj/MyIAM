package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAPIServer struct {
	*grpc.Server
	address string
}

func (gs *GrpcAPIServer) Run() {
	listen, err := net.Listen("tcp", gs.address)
	if err != nil {
		log.Fatal("failed to listen address: %s, err: %s", gs.address, err.Error())
	}
	go func() {
		err = gs.Serve(listen)
		if err != nil {
			log.Fatal("failed to serve grpc, err: %s", err.Error())
		}
	}()

}
