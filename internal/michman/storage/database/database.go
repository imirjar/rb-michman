package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/imirjar/Michman/internal/michman/models"
	_ "modernc.org/sqlite"
)

type Storage struct {
	dbConn *sql.Conn
}

func NewStorage() *Storage {
	db, err := sql.Open("sqlite", "db/divers")
	if err != nil {
		panic(err)
	}

	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}

	return &Storage{
		dbConn: conn,
	}
}

func (s Storage) GetDivers(ctx context.Context) ([]models.Diver, error) {

	rows, err := s.dbConn.QueryContext(ctx, "SELECT * FROM divers;")
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

	row := s.dbConn.QueryRowContext(ctx, "SELECT * FROM divers WHERE id=$1;", id)

	if err := row.Scan(&diver.Id, &diver.Name, &diver.Addr); err != nil {
		// log.Print("Scan error")
		return diver, err
	}

	return diver, nil
}
