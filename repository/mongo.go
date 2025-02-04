package repository

import (
	"RMazeE-server/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
)

type MongoRepository struct {
	client *mongo.Client
}

func NewMongoRepository() *MongoRepository {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. ")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &MongoRepository{
		client: client,
	}
}

func (m *MongoRepository) FindByParams(algorithm string, mazeLevel int64) ([]*types.Rank, error) {
	coll := m.client.Database("RMazeE").Collection("ranking")
	cursor, err := coll.Find(context.TODO(), bson.D{
		{"algorithm", algorithm},
		{"mazeLevel", mazeLevel},
	})
	if err != nil {
		panic(err)
	}

	results := make([]*types.Rank, 0)
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func (m *MongoRepository) mongodbTest() {
	coll := m.client.Database("RMazeE").Collection("ranking")
	name := "lopyad"
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"Name", name}}).
		Decode(&result)
	if errors.Is(err, mongo.ErrNoDocuments) {
		fmt.Printf("No document was found with the title %s\n", name)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

func (m *MongoRepository) mongodbTest2() {
	coll := m.client.Database("RMazeE").Collection("ranking")
	algorithm := "huntAndKill"
	mazeLevel := 4
	cursor, err := coll.Find(context.TODO(), bson.D{
		{"algorithm", algorithm},
		{"mazeLevel", mazeLevel},
	})
	if err != nil {
		panic(err)
	}

	var results []types.Rank
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	fmt.Println(results)
}

//func (m *MongoRepository) FindOneByName(name string) (*types.Rank, error) {
//	coll := m.client.Database("RMazeE").Collection("ranking")
//	filter := bson.D{{"name", name}}
//
//	var result types.Rank
//	err := coll.FindOne(context.TODO(), filter).Decode(&result)
//	if err != nil {
//		if errors.Is(err, mongo.ErrNoDocuments) {
//			return nil, err
//		}
//		panic(err)
//	}
//	return &result, nil
//}
