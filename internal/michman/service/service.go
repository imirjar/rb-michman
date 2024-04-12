package service

import (
	"context"
)

type Storage interface {
	GetQuery(context.Context, string) (string, error)
	ExecuteQuery(context.Context, string) (string, error)
}

type Service struct {
	storage Storage
}

func NewService() *Service {
	return &Service{
		// storage: storage.NewStorage(),
	}
}
