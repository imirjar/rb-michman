package grazer

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/storage/collector"
)

type Storage interface {
	GetDivers(context.Context) ([]models.Diver, error)
}

// Grazer must know all about active divers
// must backup information about existing divers when michan shutting down
// must ping and being pinged by divers
type Service struct {
	storage Storage
}

func New() *Service {
	return &Service{
		storage: collector.New(),
	}
}

func (s Service) DiverList(ctx context.Context) ([]models.Diver, error) {
	return s.storage.GetDivers(ctx)

}

func (s Service) LoadConnections() {} // read all connected divers, ping it, connect which is still alive

func (s Service) BackupConnections() {}
