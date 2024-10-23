package grazer

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
)

type Storage interface {
	GetDivers(context.Context) ([]models.Diver, error)
	AddDiver(context.Context, models.Diver) error
}

// Grazer must know all about active divers
// must backup information about existing divers when michan shutting down
// must ping and being pinged by divers
type Service struct {
	Storage Storage
}

func New() *Service {
	return &Service{}
}

func (s Service) DiverList(ctx context.Context) ([]models.Diver, error) {
	return s.Storage.GetDivers(ctx)
}

func (s Service) LoadConnections() {} // read all connected divers, ping it, connect which is still alive

func (s Service) BackupConnections() {}

func (s Service) ConnectDiver(ctx context.Context, diver models.Diver) error {
	return s.Storage.AddDiver(ctx, diver)
}
