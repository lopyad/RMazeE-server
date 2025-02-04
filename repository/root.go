package repository

import (
	"context"
	"log"
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	mongo *MongoRepository
	Rank  *RankRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			mongo: NewMongoRepository(),
		}
		repositoryInstance.Rank = NewRankRepository(repositoryInstance.mongo)
	})
	repositoryInstance.mongo.mongodbTest2()
	return repositoryInstance
}

func (r *Repository) GracefulShutdown() {
	log.Println("MongoClient disconnect from srv")
	if err := r.mongo.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
