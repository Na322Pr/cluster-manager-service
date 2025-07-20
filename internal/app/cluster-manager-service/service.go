package cluster_manager_service

import (
	"github.com/Na322Pr/cluster-manager-service/internal/service"
	desc "github.com/Na322Pr/cluster-manager-service/pkg/api"
)

type Implementation struct {
	desc.UnimplementedClusterManagerServer

	cmService *service.ClusterManagerService
}

func NewImplementation(cmService *service.ClusterManagerService) *Implementation {
	return &Implementation{
		cmService: cmService,
	}
}
