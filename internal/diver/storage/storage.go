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

type Storage struct {
	config       Config
	ReportsStore *reports.ReportsStore
	Target       *target.TargetDB
}

func NewStorage() *Storage {
	store := Storage{
		config: config.NewConfig(),
	}
	store.ReportsStore = reports.NewReportStore()
	store.Target = target.NewTargetDB(store.config.GetDiverTargetDB())
	return &store
}

func (s Storage) GetQuery(ctx context.Context, id string) (string, error) {
	return s.ReportsStore.GetQuery(ctx, id)
}

func (s Storage) ExecuteQuery(ctx context.Context, query string) (string, error) {
	return s.Target.SELECT(ctx, query)
}

func (s Storage) GetAllReports(ctx context.Context) (string, error) {
	return s.ReportsStore.GetAllReports(ctx)
}
