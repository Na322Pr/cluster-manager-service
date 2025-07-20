package model

import (
	"github.com/Na322Pr/kv-storage-service/pkg/nodemodel"
	"github.com/google/uuid"
	"sync"
)

type Election struct {
	selfUUID           string
	selfAddress        string
	mu                 sync.Mutex
	nodes              map[string]*nodemodel.Node
	leaderUUID         string
	electionInProgress bool
}

func NewElection(selfAddress string) *Election {
	return &Election{
		selfUUID:    uuid.New().String(),
		selfAddress: selfAddress,
		mu:          sync.Mutex{},
		nodes:       make(map[string]*nodemodel.Node),
	}
}
