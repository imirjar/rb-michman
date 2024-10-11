package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/imirjar/Michman/internal/michman/models"
	_ "modernc.org/sqlite"
)

type Storage struct {
	dbConn *sql.Conn
}

func New() *Storage {
	db, err := sql.Open("sqlite", "db/divers")
	if err != nil {
		panic(err)
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}

	store := Storage{
		dbConn: conn,
	}

	if err := store.Ping(); err != nil {
		panic(err)
	}

	return &store
}

func (s *Storage) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return s.dbConn.PingContext(ctx)
}

func (s Storage) GetDiver(ctx context.Context, id string) (models.Diver, error) {
	var diver models.Diver

	row := s.dbConn.QueryRowContext(ctx, "SELECT * FROM divers WHERE id=$1;", id)

	if err := row.Scan(&diver.Name, &diver.Addr); err != nil {
		// log.Print("Scan error")
		return diver, err
	}

	return diver, nil
}
