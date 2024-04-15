package storage

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage/api"
	"github.com/imirjar/Michman/internal/michman/storage/database"
)

type API interface {
	GetDiverReports(context.Context, string) (string, error)
}

type DB interface {
	GetDivers(context.Context) (string, error)
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
