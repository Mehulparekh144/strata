package strata

import (
	"context"
	"strata/api/strata/api"
	"strata/pkg/ds"
)

func (s *Server) LPush(ctx context.Context, req *api.ListPushRequest) (*api.ListPushResponse, error) {
	length, err := ds.LPush(s.engine, req.Key, req.Value)
	if err != nil {
		return &api.ListPushResponse{
			Length: 0,
		}, err
	}
	return &api.ListPushResponse{
		Length: length,
	}, nil
}

func (s *Server) LPop(ctx context.Context, req *api.ListPopRequest) (*api.ListPopResponse, error) {
	value, err := ds.LPop(s.engine, req.Key)
	if err != nil {
		return &api.ListPopResponse{
			Value:   "",
			Success: false,
		}, err
	}
	return &api.ListPopResponse{
		Value:   value,
		Success: true,
	}, nil
}

func (s *Server) RPush(ctx context.Context, req *api.ListPushRequest) (*api.ListPushResponse, error) {
	length, err := ds.RPush(s.engine, req.Key, req.Value)
	if err != nil {
		return &api.ListPushResponse{
			Length: 0,
		}, err
	}
	return &api.ListPushResponse{
		Length: length,
	}, nil
}

func (s *Server) RPop(ctx context.Context, req *api.ListPopRequest) (*api.ListPopResponse, error) {
	value, err := ds.RPop(s.engine, req.Key)
	if err != nil {
		return &api.ListPopResponse{
			Value:   "",
			Success: false,
		}, err
	}
	return &api.ListPopResponse{
		Value:   value,
		Success: true,
	}, nil
}
