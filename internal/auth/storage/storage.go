package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {
	db, err := sql.Open("sqlite", "users")
	if err != nil {
		panic(err)
	}

	store := Storage{
		db: db,
	}

	if err := store.Ping(); err != nil {
		panic(err)
	}

	return &store
}

func (r *Storage) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return r.db.PingContext(ctx)
}

func (r *Storage) GetUserID(ctx context.Context, username string) (int, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	var data int
	row := conn.QueryRowContext(ctx, "SELECT id FROM users WHERE username=$1;", fmt.Sprint(username))
	err = row.Scan(&data)
	if err != nil {
		return -1, err
	}

	return data, nil
}

func (r *Storage) AddUser(ctx context.Context, username, password string) (int, error) {
	return 0, nil
}
