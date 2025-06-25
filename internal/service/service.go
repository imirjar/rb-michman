package service

import (
	"context"
	"log"

	"github.com/imirjar/rb-michman/internal/models"
)

type MQ interface {
	ExecuteReport(context.Context, string) (models.Data, error)
}

type ReportsStorage interface {
	GetReports(ctx context.Context) ([]models.Report, error)
}

type Service struct {
	MQ MQ
	RS ReportsStorage
}

func New() *Service {
	return &Service{}
}

func (s Service) GetReports(ctx context.Context) ([]models.Report, error) {
	report, err := s.RS.GetReports(ctx)
	if err != nil {
		log.Print("ERROR DiverReports", err)
		return nil, err
	}
	return report, nil
	// return []models.Report{}, nil
}

func (s Service) ExecuteReport(ctx context.Context, id string) (models.Data, error) {
	return s.MQ.ExecuteReport(ctx, id)
	// return models.Data{
	// 	Columns: []string{id},
	// }, nil
}

func (s Service) ExecuteReportMap(ctx context.Context, addr, repID string) ([]map[string]interface{}, error) {
	// return s.Divers.ExecuteDiverReportMap(ctx, addr, repID)
	return []map[string]interface{}{}, nil
}
