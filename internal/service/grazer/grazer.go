package grazer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/imirjar/rb-michman/internal/models"
)

type Grazer interface {
	GetDivers(context.Context) (map[string]models.Diver, error)
	GetDiver(context.Context, string) (*models.Diver, error)
	AddDiver(context.Context, string, models.Diver) error
}

// Grazer must know all about active divers
// must backup information about existing divers when michan shutting down
// must ping and being pinged by divers
type Service struct {
	Grazer Grazer
}

func New() *Service {
	return &Service{}
}

func (s Service) DiverList(ctx context.Context) (map[string]models.Diver, error) {
	return s.Grazer.GetDivers(ctx)
}

func (s Service) DiverAddr(ctx context.Context, hash string) (string, error) {
	diver, err := s.Grazer.GetDiver(ctx, hash)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return diver.Addr, nil
}

func (s Service) ConnectDiver(ctx context.Context, diver models.Diver) error {
	bytes, err := json.Marshal(diver)
	if err != nil {
		log.Print(err)
		return err
	}

	hash := sha256.Sum256([]byte(bytes))
	return s.Grazer.AddDiver(ctx, hex.EncodeToString(hash[:]), diver)
}
