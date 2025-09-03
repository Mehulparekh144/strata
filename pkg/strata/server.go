package strata

import (
	"strata/api/strata/api"
	"strata/pkg/engine"
)

type Server struct {
	api.UnimplementedStrataServer
	engine engine.StorageEngine
}

func NewServer(e engine.StorageEngine) *Server {
	return &Server{
		engine: e,
	}
}
