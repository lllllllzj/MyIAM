package apiserver

import (
	grpcapiserver "MyIAM/internal/pkg/server"
	httpapiserver "MyIAM/internal/pkg/server"
	"MyIAM/internal/pkg/shutdown"
	"log"
)

type apiServer struct {
	httpAPIServer *httpapiserver.HttpAPIServer
	grpcAPIServer *grpcapiserver.GrpcAPIServer
	gs            *shutdown.GracefulShutdown
}

func (s *apiServer) PrepareRun() preparedApiServer {
	initRouter(s.httpAPIServer.Engine)

	return preparedApiServer{s}
}

func createAPIServer() (*apiServer, error) {
	server := &apiServer{}
	return server, nil
}

type preparedApiServer struct {
	*apiServer
}

func (ps preparedApiServer) Run() error {
	// start grpc api server
	go ps.grpcAPIServer.Run()
	//start shutdown server
	err := ps.gs.Start()
	if err != nil {
		log.Fatal()
	}
	// start gin http api server
	return ps.httpAPIServer.Run()
}
