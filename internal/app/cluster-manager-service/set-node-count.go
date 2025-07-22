package cluster_manager_service

import (
	"context"
	desc "github.com/Na322Pr/cluster-manager-service/pkg/api"
)

func (s *Implementation) SetNodeCount(
	ctx context.Context,
	req *desc.SetNodeCountRequest,
) (*desc.SetNodeCountResponse, error) {
	err := s.cmService.SetClusterSize(ctx, int(req.NodeCount))
	if err != nil {
		return &desc.SetNodeCountResponse{}, err
	}

	return &desc.SetNodeCountResponse{
		Result: true,
	}, nil
}
