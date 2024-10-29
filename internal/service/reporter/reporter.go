package reporter

import (
	"context"
	"log"

	"github.com/imirjar/rb-michman/internal/models"
)

type IDiver interface {
	// GetDiver(context.Context, string) (models.Diver, error)
	GetDiverReports(context.Context, string) ([]models.Report, error)
	ExecuteDiverReport(context.Context, string, string) (models.Report, error)
}

type Collector interface {
	GetDiver(context.Context, string) (*models.Diver, error)
}

type Service struct {
	Collector Collector
	Divers    IDiver
}

func New() *Service {
	return &Service{}
}

func (s Service) DiverReports(ctx context.Context, hash string) ([]models.Report, error) {
	report, err := s.Collector.GetDiver(ctx, hash)
	if err != nil {
		log.Print("ERROR DiverReports", err)
		return nil, err
	}
	return s.Divers.GetDiverReports(ctx, report.Addr)
}

func (s Service) GetDiverReportData(ctx context.Context, diver models.Diver) (models.Report, error) {
	return s.Divers.ExecuteDiverReport(ctx, diver.Addr, diver.Reports[0].Id)
}
