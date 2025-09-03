package strata

import (
	"context"
	"strata/api/strata/api"
	"strata/pkg/ds"
)


func (s *Server) Set(ctx context.Context, req *api.SetRequest) (*api.SetResponse, error) {
	if err := ds.Set(s.engine, req.Key, req.Value); err != nil {
		return &api.SetResponse{
			Success: false,
		}, err
	}
	return &api.SetResponse{
		Success: true,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	value, found, err := ds.Get(s.engine, req.Key)
	if err != nil {
		return &api.GetResponse{
			Value: "",
			Found: false,
		}, err
	}

	return &api.GetResponse{
		Value: value,
		Found: found,
	}, nil
}

func (s *Server) Del(ctx context.Context, req *api.DelRequest) (*api.DelResponse, error) {
	found, err := ds.Del(s.engine, req.Key)
	if err != nil {
		return &api.DelResponse{
			Success: false,
		}, err
	}
	return &api.DelResponse{
		Success: found,
	}, nil
}