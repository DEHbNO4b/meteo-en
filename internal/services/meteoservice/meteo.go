package meteoservice

import (
	"context"
	"meteo-lightning/internal/domain/models"
)

type MeteoStore interface {
	SaveMeteoData(ctx context.Context, data []models.MeteoData) error
}
type MeteoService struct {
	meteoDB MeteoStore
}

func NewService(db MeteoStore) *MeteoService {
	return &MeteoService{meteoDB: db}
}

func (ms *MeteoService) SaveMeteoData(ctx context.Context, data []models.MeteoData) error {
	return ms.meteoDB.SaveMeteoData(ctx, data)
}
