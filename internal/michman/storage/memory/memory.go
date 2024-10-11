package memory

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
)

type Storage struct {
	divers map[string]models.Diver
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) GetDivers(ctx context.Context) (map[string]models.Diver, error) {
	return s.divers, nil
}

func (s *Storage) AddDiver([]models.Diver) (int, error) {
	return 1, nil
}
