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

func (gas *GrpcAPIServer) Run() {
	listen, err := net.Listen("tcp", gas.address)
	if err != nil {
		log.Fatal("failed to listen address: %s, err: %s", gas.address, err.Error())
	}
	go func() {
		err = gas.Serve(listen)
		if err != nil {
			log.Fatal("failed to serve grpc, err: %s", err.Error())
		}
	}()

}
