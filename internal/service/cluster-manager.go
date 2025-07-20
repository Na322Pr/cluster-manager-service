package service

import (
	"context"
	"github.com/Na322Pr/cluster-manager-service/internal/model"
	"go.uber.org/zap"

	nomad "github.com/hashicorp/nomad/api"
)

type ClusterManagerService struct {
	cluster  *model.Cluster
	election *model.Election

	nomadClient *nomad.Client
	logger      *zap.Logger
}

func NewClusterManagerService(selfAddress string, nomadClient *nomad.Client, logger *zap.Logger) *ClusterManagerService {
	return &ClusterManagerService{
		cluster:     model.NewSampleCluster(),
		election:    model.NewElection(selfAddress),
		logger:      logger,
		nomadClient: nomadClient,
	}
}

func (s *ClusterManagerService) RunCluster() error {
	nomadJob := s.cluster.GetNomadJob()

	_, _, err := s.nomadClient.Jobs().Register(nomadJob, nil)
	return err
}

func (s *ClusterManagerService) SetClusterSize(_ context.Context, size int) error {
	s.cluster.SetClusterSize(size)

	_, _, err := s.nomadClient.Jobs().Scale(
		s.cluster.GetJobID(),
		s.cluster.GetTaskGroupName(),
		s.cluster.GetClusterSize(),
		"",
		false,
		nil,
		nil,
	)

	return err
}
