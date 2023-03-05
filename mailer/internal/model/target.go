package model

type Target struct {
	ID    string            `json:"id"`
	Email string            `json:"email"`
	Model map[string]string `json:"model"`
}

type TargetRepository struct {
	*Repository[Target]
}

func NewTargetRepository() *TargetRepository {
	return &TargetRepository{
		Repository: &Repository[Target]{
			name:  "targets",
			entry: new(Target),
		},
	}
}
