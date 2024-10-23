package memory

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/models"
)

type Storage struct {
	divers []models.Diver
}

func New() *Storage {
	store := Storage{
		divers: []models.Diver{
			models.Diver{
				Name: "diver1",
				Addr: "192.168.0.1",
			},
			models.Diver{
				Name: "diver2",
				Addr: "192.168.0.1",
			},
			models.Diver{
				Name: "diver3",
				Addr: "192.168.0.1",
			},
		},
	}
	return &store
}

func (s *Storage) GetDivers(ctx context.Context) ([]models.Diver, error) {
	return s.divers, nil
}

func (s *Storage) AddDiver(ctx context.Context, diver models.Diver) error {
	s.divers = append(s.divers, diver)
	return nil
}
