package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRepository[T any] interface {
	Save(t T) string
	FindOne(id string) *T
	UpdateByID(id string, t T) int64
	DeleteOne(id string) int64
}

type Entity struct {
	ID primitive.ObjectID `json:"_id"`
}

type Repository[T any] struct {
	Name       string
	connection *mongo.Database
}

func (repo *Repository[T]) Save(t T) int64 {
	collection := repo.connection.Collection(repo.Name)
	_, err := collection.InsertOne(context.Background(), t)
	if err != nil {
		return 0
	}
	return 1
}

func (repo *Repository[T]) FindOne(id string) *T {
	collection := repo.connection.Collection(repo.Name)
	var t T
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&t)
	if err != nil {
		panic(err)
	}
	return &t
}

func (repo *Repository[T]) UpdateByID(id string, t T) int64 {
	collection := repo.connection.Collection(repo.Name)
	updated, err := collection.UpdateByID(context.Background(), bson.M{"_id": id}, t)
	if err != nil {
		return 0
	}
	return updated.ModifiedCount
}

func (repo *Repository[T]) DeleteOne(id string) int64 {
	collection := repo.connection.Collection(repo.Name)
	updated, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return 0
	}
	return updated.DeletedCount
}
