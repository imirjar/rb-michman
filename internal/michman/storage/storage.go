package storage

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage/api"
	"github.com/imirjar/Michman/internal/michman/storage/database"
)

type API interface {
	GetDiverReports(context.Context, string) ([]models.Report, error)
	ExecuteDiverReport(context.Context, string, string) (models.Report, error)
}

type DB interface {
	GetDivers(context.Context) ([]models.Diver, error)
	GetDiver(context.Context, string) (models.Diver, error)
}

type Storage struct {
	DB
	API
}

func NewStorage() *Storage {
	return &Storage{
		DB:  database.NewStorage(),
		API: api.NewAPI(),
	}
}
