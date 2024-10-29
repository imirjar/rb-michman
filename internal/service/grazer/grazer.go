package grazer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/imirjar/rb-michman/internal/models"
)

type Collector interface {
	GetDivers(context.Context) (map[string]models.Diver, error)
	AddDiver(context.Context, string, models.Diver) error
}

type IDiver interface {
	CheckConnection(context.Context, models.Diver) bool
}

// Grazer must know all about active divers
// must backup information about existing divers when michan shutting down
// must ping and being pinged by divers
type Service struct {
	Collector Collector
	Divers    IDiver
}

func New() *Service {
	return &Service{}
}

func (s Service) DiverList(ctx context.Context) (map[string]models.Diver, error) {
	return s.Collector.GetDivers(ctx)
}

func (s Service) CheckConnections(ctx context.Context) error {
	divers, err := s.Collector.GetDivers(ctx)
	if err != nil {
		log.Print(err)
		return err
	}
	for _, dvr := range divers {
		dvr.Connected = s.Divers.CheckConnection(ctx, dvr)
	}
	return nil
} // read all connected divers, ping it, connect which is still alive

func (s Service) ConnectDiver(ctx context.Context, diver models.Diver) error {
	bytes, err := json.Marshal(diver)
	if err != nil {
		log.Print(err)
		return err
	}

	hash := sha256.Sum256([]byte(bytes))
	return s.Collector.AddDiver(ctx, hex.EncodeToString(hash[:]), diver)
}
