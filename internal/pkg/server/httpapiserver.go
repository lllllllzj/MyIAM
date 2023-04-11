package server

import "github.com/gin-gonic/gin"

type HttpAPIServer struct {
	*gin.Engine
}

func (hs *HttpAPIServer) Run() error {
	return nil
}
