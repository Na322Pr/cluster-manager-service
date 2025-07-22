package service

import (
	"context"
	"github.com/Na322Pr/cluster-manager-service/internal/model"
	"go.uber.org/zap"

	consul "github.com/hashicorp/consul/api"
	nomad "github.com/hashicorp/nomad/api"
)

type ClusterManagerService struct {
	cluster  *model.Cluster
	election *model.Election

	consulClient *consul.Client
	nomadClient  *nomad.Client
	logger       *zap.Logger
}

func NewClusterManagerService(
	selfAddress string,
	consulClient *consul.Client,
	nomadClient *nomad.Client,
	logger *zap.Logger,
) *ClusterManagerService {
	return &ClusterManagerService{
		cluster:      model.NewSampleCluster(),
		election:     model.NewElection(selfAddress),
		consulClient: consulClient,
		nomadClient:  nomadClient,
		logger:       logger,
	}
}

func (s *ClusterManagerService) RunCluster() error {
	nomadJob := s.cluster.GetNomadJob()

	_, _, err := s.nomadClient.Jobs().Register(nomadJob, nil)
	return err
}

func (s *ClusterManagerService) SetClusterSize(_ context.Context, size int) error {
	s.cluster.SetClusterSize(size)

	meta := make(map[string]interface{})
	
	_, _, err := s.nomadClient.Jobs().Scale(
		s.cluster.GetJobID(),
		s.cluster.GetTaskGroupName(),
		s.cluster.GetClusterSize(),
		"",
		false,
		meta,
		nil,
	)

	return err
}
