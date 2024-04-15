package storage

import (
	"context"

	"github.com/imirjar/Michman/internal/diver/storage/reports"
	"github.com/imirjar/Michman/internal/diver/storage/target"
)

// mongo db with reports
// must connect sqlite db file
type ReportsStore interface {
	GetQuery(context.Context, string) (string, error)
	GetAllReports(context.Context) (string, error)
}

// target db where reports from mongo are sending
// must connect db by env file
type Target interface {
	SELECT(context.Context, string) (string, error)
}

type Storage struct {
	ReportsStore
	Target
}

func NewStorage() *Storage {
	return &Storage{
		ReportsStore: reports.NewReportStore(),
		Target:       target.NewTargetDB(),
	}
}

func (s Storage) GetQuery(ctx context.Context, id string) (string, error) {
	return s.ReportsStore.GetQuery(ctx, id)
}

func (s Storage) ExecuteQuery(ctx context.Context, query string) (string, error) {
	return s.Target.SELECT(ctx, query)
}

// CREATE TABLE IF NOT EXISTS metrics (
// 	id varchar NOT NULL,
// 	"type" varchar NOT NULL,s
// 	value float8 NULL,
// 	CONSTRAINT metrics_pk PRIMARY KEY (id)
// );
