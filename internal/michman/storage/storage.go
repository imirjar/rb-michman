package storage

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage/api"
	"github.com/imirjar/Michman/internal/michman/storage/memory"
)

type API interface {
	GetDiverReports(context.Context, string) ([]models.Report, error)
	ExecuteDiverReport(context.Context, string, string) (models.Report, error)
}

type Memory interface {
	GetDivers(ctx context.Context) (map[string]models.Diver, error)
	AddDiver([]models.Diver) (int, error)
}

type Storage struct {
	Memory
	API
}

func New() *Storage {
	return &Storage{
		Memory: memory.New(),
		API:    api.New(),
	}
}
