package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository[T any] interface {
	Save(t T) int64
	FindOne(id string) *T
	FindBy(attr, value string) []T
	FinAll() []T
	Paginate(page, limit int64) []T
	Count() int64
	UpdateByID(id string, t T) int64
	DeleteOne(id string) int64
}

type Repository[T any] struct {
	Name       string
	connection *mongo.Database
	context    context.Context
}

func (repo *Repository[T]) Save(ts []T) int64 {
	collection := repo.connection.Collection(repo.Name)
	var count int64 = 0
	for _, t := range ts {
		_, err := collection.InsertOne(repo.context, t)
		if err != nil {
			count++
		}
	}
	return count
}

func (repo *Repository[T]) Count() int64 {
	collection := repo.connection.Collection(repo.Name)
	count, err := collection.CountDocuments(repo.context, bson.M{}, nil)
	if err != nil {
		return -1
	}
	return count
}

func (repo *Repository[T]) FinAll() []T {
	return repo.Paginate(1, repo.Count())
}

func (repo *Repository[T]) Paginate(page, limit int64) []T {
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

func (repo *Repository[T]) FindOne(id string) *T {
	collection := repo.connection.Collection(repo.Name)
	var t T
	err := collection.FindOne(repo.context, bson.M{"_id": id}).Decode(&t)
	if err != nil {
		panic(err)
	}
	return &t
}

func (repo *Repository[T]) FindBy(attr, value string) []T {
	collection := repo.connection.Collection(repo.Name)
	filter := bson.M{attr: value}

	cursor, err := collection.Find(repo.context, filter)
	if err != nil {
		return nil
	}

	defer cursor.Close(repo.context)

	var results []T
	if err2 := cursor.All(repo.context, &results); err2 != nil {
		return nil
	}
	return results
}

func (repo *Repository[T]) UpdateByID(id string, t T) int64 {
	collection := repo.connection.Collection(repo.Name)
	updated, err := collection.UpdateByID(repo.context, bson.M{"_id": id}, t)
	if err != nil {
		return 0
	}
	return updated.ModifiedCount
}

func (repo *Repository[T]) DeleteOne(id string) int64 {
	collection := repo.connection.Collection(repo.Name)
	updated, err := collection.DeleteOne(repo.context, bson.M{"_id": id})
	if err != nil {
		return 0
	}
	return updated.DeletedCount
}
