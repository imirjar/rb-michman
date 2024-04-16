package service

import (
	"context"
	"fmt"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage"
)

type Storage interface {
	GetDivers(context.Context) (string, error)
	GetDiver(context.Context, string) (models.Diver, error)
	GetDiverReports(context.Context, string) (string, error)
}

type Service struct {
	storage Storage
}

func (s Service) DiverList(ctx context.Context) (string, error) {
	return s.storage.GetDivers(ctx)
}

func (s Service) DiverInfo(ctx context.Context, id string) (string, error) {

	diver, err := s.storage.GetDiver(ctx, id)
	if err != nil {
		return err.Error(), err
	}

	diver.Reports, err = s.storage.GetDiverReports(ctx, diver.Addr)
	if err != nil {
		return err.Error(), err
	}

	return fmt.Sprint(diver), err
}

func NewService() *Service {
	return &Service{
		storage: storage.NewStorage(),
	}
}
