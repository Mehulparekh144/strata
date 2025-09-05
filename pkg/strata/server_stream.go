package strata

import (
	"strata/api/strata/api"
	"strata/pkg/ds"
	"strings"

	"google.golang.org/grpc"
)

func (s *Server) XReadStream(req *api.StreamRequest, stream grpc.ServerStreamingServer[api.StreamResponse]) error {
	for event := range ds.EventBus {
		if req.MutationKey == "" || strings.HasPrefix(event.Key, req.MutationKey) {
			if err := stream.Send(event); err != nil {
				return err
			}
		}
	}
	return nil
}
