package service

import (
	"RMazeE-server/repository"
	"sync"
)

//network, repository 의 가교

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository

	Rank *Rank
}

func NewService(repo *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: repo,
		}

		serviceInstance.Rank = newRankService(repo.Rank)
	})

	return serviceInstance
}

func (s *Service) GracefulShutdown() {
	s.repository.GracefulShutdown()
}
