package reporter

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage/diver"
)

type Storager interface {
	// GetDiver(context.Context, string) (models.Diver, error)
	GetDiverReports(context.Context, string) ([]models.Report, error)
	ExecuteDiverReport(context.Context, string, string) (models.Report, error)
}

type MemStorager interface {
	GetDiver(context.Context, string) (models.Diver, error)
}

type Service struct {
	storage    Storager
	memStorage MemStorager
}

func New() *Service {
	return &Service{
		storage: diver.New(),
	}
}

func (s Service) DiverReports(ctx context.Context, id string) ([]models.Report, error) {
	diver, err := s.memStorage.GetDiver(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.storage.GetDiverReports(ctx, diver.Addr)
}

func (s Service) GetDiverReportData(ctx context.Context, diver models.Diver) (models.Report, error) {

	return s.storage.ExecuteDiverReport(ctx, diver.Addr, diver.Reports[0].Id)
}
