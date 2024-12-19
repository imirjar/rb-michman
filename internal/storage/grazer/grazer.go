package collector

import (
	"context"

	"github.com/imirjar/rb-michman/internal/models"
)

type Storage struct {
	divers map[string]models.Diver
}

func New() *Storage {
	store := &Storage{
		divers: map[string]models.Diver{},
	}

	return store
}

func (s *Storage) GetDivers(ctx context.Context) (map[string]models.Diver, error) {
	return s.divers, nil
}

func (s *Storage) AddDiver(ctx context.Context, hash string, diver models.Diver) error {
	s.divers[hash] = diver
	return nil
}

func (s *Storage) GetDiver(ctx context.Context, hash string) (*models.Diver, error) {
	diver := s.divers[hash]
	return &diver, nil
}
