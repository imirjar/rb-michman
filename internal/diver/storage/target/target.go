package target

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TargetDB struct {
	db *pgx.Conn
}

func NewTargetDB(dbConn string) *TargetDB {
	conn, err := pgx.Connect(context.Background(), dbConn)
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
