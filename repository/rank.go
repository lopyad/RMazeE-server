package repository

import (
	"RMazeE-server/types"
	"RMazeE-server/types/errors"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type RankRepository struct {
	Mongo   *MongoRepository
	rankMap []*types.Rank
}

func NewRankRepository(mongo *MongoRepository) *RankRepository {
	return &RankRepository{
		Mongo:   mongo,
		rankMap: []*types.Rank{},
	}
}

func (u *RankRepository) Create(newRanking []*types.Rank) error {
	coll := u.Mongo.client.Database("RMazeE").Collection("ranking")

	_, err := coll.InsertMany(context.TODO(), newRanking)
	return err
}

func (u *RankRepository) Get(algorithm string, mazeLevel int64) ([]*types.Rank, error) {
	return u.Mongo.FindByParams(algorithm, mazeLevel)
}

func (u *RankRepository) Update(newRank *types.Rank) error {
	var isExisted bool = false
	for _, curRank := range u.rankMap {
		if newRank.Name == curRank.Name {
			curRank.ElapsedTime = newRank.ElapsedTime
			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *RankRepository) DeleteMany(algorithm string, mazeLevel int64) error {
	coll := u.Mongo.client.Database("RMazeE").Collection("ranking")
	filter := bson.D{
		{"algorithm", algorithm},
		{"mazeLevel", mazeLevel},
	}
	_, err := coll.DeleteMany(context.TODO(), filter)
	return err
}
