package apiserver

import "github.com/gin-gonic/gin"
import grpcapiserver "internal/pkg/server"

type apiServer struct {
	g             *gin.Engine
	grpcAPIServer *grpcapiserver.GrpcAPIServer
}

func (as *apiServer) PrepareRun() {
	initRouter(as.g)
}

type preparedApiServer struct {
	*apiServer
}

func (pas *preparedApiServer) Run() {
	go pas.grpcAPIServer.Run()
	err := g.
}
