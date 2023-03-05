package model

import "github.com/go-redis/redis"

type Target struct {
	ID    string            `json:"id"`
	Email string            `json:"email"`
	Model map[string]string `json:"model"`
}

type TargetRepository struct {
	*Repository[Target]
}

func NewTargetRepository(conn *redis.Client) *TargetRepository {
	return &TargetRepository{
		Repository: &Repository[Target]{
			conn:  conn,
			name:  "targets",
			entry: new(Target),
		},
	}
}
