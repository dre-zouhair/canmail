package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoRepository[T any] interface {
	Save(t T) string
	FindOne(id string) *T
	FinAll() []T
	Paginate(page, limit int64) []T
	Count() int64
	UpdateByID(id string, t T) int64
	DeleteOne(id string) int64
}

type Entity struct {
	ID primitive.ObjectID `json:"_id"`
}

type MongoRepository[T any] struct {
	Name       string
	connection *mongo.Database
	context    context.Context
}

func (repo *MongoRepository[T]) Save(t T) int64 {
	collection := repo.connection.Collection(repo.Name)
	_, err := collection.InsertOne(repo.context, t)
	if err != nil {
		return 0
	}
	return 1
}

func (repo *MongoRepository[T]) Count() int64 {
	collection := repo.connection.Collection(repo.Name)
	count, err := collection.CountDocuments(repo.context, bson.M{}, nil)
	if err != nil {
		return -1
	}
	return count
}

func (repo *MongoRepository[T]) FinAll() []T {
	return repo.Paginate(1, repo.Count())
}

func (repo *MongoRepository[T]) Paginate(page, limit int64) []T {
	collection := repo.connection.Collection(repo.Name)

	findOptions := options.Find().SetSkip((page - 1) * limit).SetLimit(limit)
	filter := bson.M{}

	cursor, err := collection.Find(repo.context, filter, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer cursor.Close(repo.context)

	var ts []T

	for cursor.Next(repo.context) {
		var t T
		_ = cursor.Decode(&t)
		ts = append(ts, t)
	}

	return ts
}

func (repo *MongoRepository[T]) FindOne(id string) *T {
	collection := repo.connection.Collection(repo.Name)
	var t T
	err := collection.FindOne(repo.context, bson.M{"_id": id}).Decode(&t)
	if err != nil {
		panic(err)
	}
	return &t
}

func (repo *MongoRepository[T]) UpdateByID(id string, t T) int64 {
	collection := repo.connection.Collection(repo.Name)
	updated, err := collection.UpdateByID(repo.context, bson.M{"_id": id}, t)
	if err != nil {
		return 0
	}
	return updated.ModifiedCount
}

func (repo *MongoRepository[T]) DeleteOne(id string) int64 {
	collection := repo.connection.Collection(repo.Name)
	updated, err := collection.DeleteOne(repo.context, bson.M{"_id": id})
	if err != nil {
		return 0
	}
	return updated.DeletedCount
}