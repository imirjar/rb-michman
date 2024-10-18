package reporter

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage/diver"
)

type IDiver interface {
	// GetDiver(context.Context, string) (models.Diver, error)
	GetDiverReports(context.Context, string) ([]models.Report, error)
	ExecuteDiverReport(context.Context, string, string) (models.Report, error)
}

type Collector interface {
	GetDiverReports(context.Context, string) ([]models.Report, error)
}

type Service struct {
	diverAPI        IDiver
	reportCollector Collector
}

func New() *Service {
	return &Service{
		diverAPI: diver.New(),
	}
}

func (s Service) DiverReports(ctx context.Context, id string) ([]models.Report, error) {
	return s.reportCollector.GetDiverReports(ctx, id)
}

func (s Service) GetDiverReportData(ctx context.Context, diver models.Diver) (models.Report, error) {
	return s.diverAPI.ExecuteDiverReport(ctx, diver.Addr, diver.Reports[0].Id)
}
