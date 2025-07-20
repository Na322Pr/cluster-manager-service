package cluster_manager_service

import (
	"context"
	desc "github.com/Na322Pr/cluster-manager-service/pkg/api"
)

func (s *Implementation) SetNodeVote(
	ctx context.Context,
	req *desc.SetNodeVoteRequest,
) (*desc.SetNodeVoteResponse, error) {
	//peers := s.nodeService.SeedHandleAddPeer(req.GetAddress())

	return &desc.SetNodeVoteResponse{
		Result: true,
	}, nil
}
