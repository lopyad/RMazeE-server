package service

import (
	"RMazeE-server/repository"
	"RMazeE-server/types"
)

type Rank struct {
	rankRepository *repository.RankRepository
}

func newRankService(rankRepository *repository.RankRepository) *Rank {
	return &Rank{rankRepository: rankRepository}
}

func (u *Rank) Create(newRank *types.Rank) error {
	if foundRanking, err := u.rankRepository.Mongo.FindByParams(newRank.Algorithm, newRank.MazeLevel); err != nil {
		return err
	} else {
		existFlag := false
		for _, curRank := range foundRanking {
			if curRank.Name == newRank.Name {
				if curRank.ElapsedTime < newRank.ElapsedTime {
					return nil
				}
				curRank.ElapsedTime = newRank.ElapsedTime
				newRank = curRank
				existFlag = true
				break
			}
		}

		if existFlag {
			u.rankRepository.DeleteMany(newRank.Algorithm, newRank.MazeLevel)
			for _, curRank := range foundRanking {
				if curRank.ElapsedTime > newRank.ElapsedTime &&
					curRank.Rank < newRank.Rank {
					newRank.Rank = min(newRank.Rank, curRank.Rank)
					curRank.Rank += 1
				}
			}
			u.rankRepository.Create(foundRanking)
			return nil
		} else {
			u.rankRepository.DeleteMany(newRank.Algorithm, newRank.MazeLevel)
			newRank.Rank = int64(len(foundRanking) + 1)
			for _, curRank := range foundRanking {
				if curRank.ElapsedTime > newRank.ElapsedTime {
					newRank.Rank = min(newRank.Rank, curRank.Rank)
					curRank.Rank += 1
				}
			}
			foundRanking = append(foundRanking, newRank)
			u.rankRepository.Create(foundRanking)
			return nil
		}

	}
}

func (u *Rank) Update(newRank *types.Rank) error {
	return u.rankRepository.Update(newRank)
}

func (u *Rank) Delete(user *types.Rank) error {
	return nil
}

func (u *Rank) Get(algorithm string, mazeLevel int64) ([]*types.Rank, error) {
	return u.rankRepository.Get(algorithm, mazeLevel)
}
