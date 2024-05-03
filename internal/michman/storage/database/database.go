package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/imirjar/Michman/internal/michman/models"
	_ "modernc.org/sqlite"
)

type Storage struct {
	driver string
	path   string
}

func NewStorage() *Storage {
	return &Storage{
		path:   "divers",
		driver: "sqlite",
	}
}

func (s Storage) GetDivers(ctx context.Context) ([]models.Diver, error) {
	db, err := sql.Open(s.driver, s.path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := conn.QueryContext(ctx, "SELECT * FROM divers;")
	if err != nil {
		return nil, err
	}

	divers := make([]models.Diver, 0)
	for rows.Next() {
		var diver models.Diver
		if err = rows.Scan(&diver.Id, &diver.Name, &diver.Addr); err != nil {
			return nil, err
		}
		divers = append(divers, diver)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return divers, nil
}

func (s Storage) GetDiver(ctx context.Context, id string) (models.Diver, error) {
	var diver models.Diver
	db, err := sql.Open(s.driver, s.path)
	if err != nil {
		return diver, err
	}
	defer db.Close()

	conn, err := db.Conn(ctx)
	if err != nil {
		// log.Print("Conn")
		return diver, err
	}

	row := conn.QueryRowContext(ctx, "SELECT * FROM divers WHERE id=$1;", id)

	if err = row.Scan(&diver.Id, &diver.Name, &diver.Addr); err != nil {
		// log.Print("Scan error")
		return diver, err
	}

	return diver, nil
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
