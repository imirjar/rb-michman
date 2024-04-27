package service

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage"
)

type Storage interface {
	GetDivers(context.Context) ([]models.Diver, error)
	GetDiver(context.Context, string) (models.Diver, error)
	GetDiverReports(context.Context, string) ([]models.Report, error)
}

type Service struct {
	storage Storage
}

func NewService() *Service {
	return &Service{
		storage: storage.NewStorage(),
	}
}

func (s Service) DiverList(ctx context.Context) ([]models.Diver, error) {
	return s.storage.GetDivers(ctx)

}

func (s Service) DiverReports(ctx context.Context, id string) ([]models.Report, error) {

	diver, err := s.storage.GetDiver(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.storage.GetDiverReports(ctx, diver.Addr)
}
