package enservice

import (
	"context"
	"meteo-lightning/internal/domain/models"
)

type ENStore interface {
	SaveEnData(ctx context.Context, data []models.StrokeEN) error
}
type ENService struct {
	enDB ENStore
}

func NewService(db ENStore) *ENService {
	return &ENService{enDB: db}
}

func (ms *ENService) SaveEnData(ctx context.Context, data []models.StrokeEN) error {
	return ms.enDB.SaveEnData(ctx, data)
}
