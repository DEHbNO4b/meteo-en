package enservice

import (
	"context"
	"meteo-lightning/internal/domain/models"
)

type ENStore interface {
	SaveEnData(ctx context.Context, data []models.StrokeEN) error
	Close()
}

type ENService struct {
	enDB ENStore
}

func NewService(db ENStore) *ENService {
	return &ENService{enDB: db}
}

func (es *ENService) SaveEnData(ctx context.Context, data []models.StrokeEN) error {
	return es.enDB.SaveEnData(ctx, data)
}

func (es *ENService) Close() {
	es.enDB.Close()
}
