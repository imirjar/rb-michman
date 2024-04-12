package storage

import (
	"context"
	"database/sql"
	"time"
)

type Storage struct {
	driver string
	path   string
}

func NewStorage() *Storage {
	return &Storage{
		path:   "reports",
		driver: "sqlite",
	}
}

func (s Storage) GetDivers(ctx context.Context, id string) (string, error) {
	db, err := sql.Open(s.driver, s.path)
	if err != nil {
		return err.Error(), err
	}
	defer db.Close()

	conn, err := db.Conn(ctx)
	if err != nil {
		return err.Error(), err
	}

	var data string

	row := conn.QueryRowContext(ctx, "SELECT query FROM reports;")
	err = row.Scan(&data)
	if err != nil {
		return err.Error(), err
	}

	return data, nil
}

func (s *Storage) Ping() error {
	db, err := sql.Open(s.driver, s.path)
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}
