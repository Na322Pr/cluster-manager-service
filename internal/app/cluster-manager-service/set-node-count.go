package cluster_manager_service

import (
	"context"
	desc "github.com/Na322Pr/cluster-manager-service/pkg/api"
)

func (s *Implementation) SetNodeCount(
	ctx context.Context,
	req *desc.SetNodeCountRequest,
) (*desc.SetNodeCountResponse, error) {
	//peers := s.nodeService.SeedHandleAddPeer(req.GetAddress())

	return &desc.SetNodeCountResponse{
		Result: true,
	}, nil
}
