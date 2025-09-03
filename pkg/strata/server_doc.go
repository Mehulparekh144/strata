package strata

import (
	"context"
	"strata/api/strata/api"
	"strata/pkg/ds"
)


func (s *Server) DocSet(ctx context.Context, req *api.DocSetRequest) (*api.DocSetResponse, error) {
	if err := ds.DocSet(s.engine, req.Key, req.Json); err != nil {
		return &api.DocSetResponse{
			Success: false,
		}, err
	}
	return &api.DocSetResponse{
		Success: true,
	}, nil
}

func (s *Server) DocGet(ctx context.Context, req *api.DocGetRequest) (*api.DocGetResponse, error) {
	json, found, err := ds.DocGet(s.engine, req.Key, req.Path)
	if err != nil {
		return &api.DocGetResponse{
			Json: "",
			Found: false,
		}, err
	}
	return &api.DocGetResponse{
		Json: json,
		Found: found,
	}, nil
}

func (s *Server) DocDel(ctx context.Context, req *api.DocDelRequest) (*api.DocDelResponse, error) {
	found, err := ds.DocDel(s.engine, req.Key)
	if err != nil {
		return &api.DocDelResponse{
			Success: false,
		}, err
	}
	return &api.DocDelResponse{
		Success: found,
	}, nil
}