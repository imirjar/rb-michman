package service

import (
	"context"
	"encoding/json"

	"github.com/imirjar/Michman/internal/diver/storage"
)

type Service struct {
	storage Storage
}

type Storage interface {
	GetQuery(context.Context, string) (string, error)
	ExecuteQuery(context.Context, string) ([]map[string]any, error)
	GetAllReports(ctx context.Context) (string, error)
}

func NewService() *Service {
	return &Service{
		storage: storage.NewStorage(),
	}
}

func (s Service) Execute(ctx context.Context, id string) (string, error) {
	report, err := s.storage.GetQuery(ctx, id)
	if err != nil {
		return "", err
	}

	data, err := s.storage.ExecuteQuery(ctx, report)
	if err != nil {
		return "", err
	}

	m, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(m), nil
}

func (s Service) ReportsList(ctx context.Context) (string, error) {
	return s.storage.GetAllReports(ctx)
}
