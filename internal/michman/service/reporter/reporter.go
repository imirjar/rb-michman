package reporter

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
)

type Diver interface {
	// GetDiver(context.Context, string) (models.Diver, error)
	GetDiverReports(context.Context, string) ([]models.Report, error)
	ExecuteDiverReport(context.Context, string, string) (models.Report, error)
}

type Reporter interface {
	GetDiverReports(context.Context, string) ([]models.Report, error)
}

type Service struct {
	DiverStore  Diver
	ReportStore Reporter
}

func New() *Service {
	return &Service{}
}

func (s Service) DiverReports(ctx context.Context, id string) ([]models.Report, error) {
	return s.ReportStore.GetDiverReports(ctx, id)
}

func (s Service) GetDiverReportData(ctx context.Context, diver models.Diver) (models.Report, error) {
	return s.DiverStore.ExecuteDiverReport(ctx, diver.Addr, diver.Reports[0].Id)
}
