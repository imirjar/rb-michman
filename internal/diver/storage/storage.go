package storage

import (
	"context"

	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/diver/storage/reports"
	"github.com/imirjar/Michman/internal/diver/storage/target"
)

type Config interface {
	GetDiverTargetDB() string
}

type ReportsStore interface {
	GetQuery(ctx context.Context, id string) (string, error)
	GetAllReports(ctx context.Context) (string, error)
}

type Target interface {
	ExecuteQuery(ctx context.Context, query string) ([]map[string]any, error)
}

type Storage struct {
	ReportsStore
	Target
}

func NewStorage() *Storage {
	var config Config = config.NewConfig()
	return &Storage{
		ReportsStore: reports.NewReportStore(),
		Target:       target.NewTargetDB(config.GetDiverTargetDB()),
	}
}
