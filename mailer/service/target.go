package service

import (
	"context"
	"github.com/dre-zouhair/mailer/internal/db"
	"github.com/dre-zouhair/mailer/internal/model"
)

type TargetService struct {
	repo *model.TargetRepository
}

func NewTargetService() *TargetService {
	connection, dbName := db.GetDBConnection(context.Background())

	TargetRepository := model.NewTargetMongoRepository(connection.Database(dbName))
	return &TargetService{
		TargetRepository,
	}
}

func (service *TargetService) AddTarget(t model.Target) int64 {
	return service.repo.Save([]model.Target{t})
}

func (service *TargetService) SaveTargets(t []model.Target) int64 {
	return service.repo.Save(t)
}

func (service *TargetService) UpdateTarget(t model.Target) int64 {
	return service.repo.UpdateByID(t.Email, t)
}

func (service *TargetService) RemoveTarget(t model.Target) int64 {
	return service.repo.DeleteOne(t.Email)
}

func (service *TargetService) Get(id string) *model.Target {
	return service.repo.FindOne(id)
}

func (service *TargetService) Retrieve(page, limit int64) []model.Target {
	return service.repo.Paginate(page, limit)
}

func (service *TargetService) RetrieveAll() []model.Target {
	return service.repo.FinAll()
}
