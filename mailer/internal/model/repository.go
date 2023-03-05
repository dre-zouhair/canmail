package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"reflect"
)

type IRepository[T any] interface {
	GetAll(conn *redis.Client) []T
	Get(conn *redis.Client, key string) *T
	Unmarshal([]byte) error
}

type Repository[T any] struct {
	conn    *redis.Client
	name    string
	entries []T
	entry   *T
}

func (entity *Repository[T]) GetAll() []T {
	result, err := entity.conn.LRange(entity.name, 0, -1).Result()
	if err != nil {
		panic(err)
	}
	entity.entries = make([]T, 0)
	for _, jsonStr := range result {
		err := entity.Unmarshal([]byte(jsonStr))
		if err != nil {
			fmt.Println(err)
			continue
		}
		entity.entries = append(entity.entries, *entity.entry)
	}
	return entity.entries
}

func (entity *Repository[T]) Get(keyName string, keyValue any) *T {
	if len(entity.entries) == 0 {
		_ = entity.GetAll()
	}

	for _, entry := range entity.entries {
		entryReflect := reflect.ValueOf(&entry)
		field := reflect.Indirect(entryReflect).FieldByName(keyName)
		valueReflect := reflect.ValueOf(keyValue)
		if field.String() == valueReflect.String() {
			return &entry
		}
	}

	return nil
}

func (entity *Repository[T]) Unmarshal(bytes []byte) error {
	return json.Unmarshal(bytes, entity.entry)
}
