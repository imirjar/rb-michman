package target

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TargetDB struct {
	db *pgx.Conn
}

func NewTargetDB() *TargetDB {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &TargetDB{
		db: conn,
	}
}

func (t *TargetDB) SELECT(ctx context.Context, query string) (string, error) {
	rows := t.db.QueryRow(ctx, query)

	var value string
	err := rows.Scan(&value)
	if err != nil {
		return err.Error(), err
	}

	return value, nil
}
