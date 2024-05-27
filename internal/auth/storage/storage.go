package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

type UserStore struct {
	dbConn *sql.Conn
}

func NewStorage() *UserStore {
	db, err := sql.Open("sqlite", "db/users")
	if err != nil {
		panic(err)
	}
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}

	store := UserStore{
		dbConn: conn,
	}

	if err := store.Ping(); err != nil {
		panic(err)
	}

	return &store
}

func (s *UserStore) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return s.dbConn.PingContext(ctx)
}

func (s *UserStore) GetUserID(ctx context.Context, username string) (int, error) {

	var data int
	row := s.dbConn.QueryRowContext(ctx, "SELECT id FROM users WHERE username=$1;", fmt.Sprint(username))
	err := row.Scan(&data)
	if err != nil {
		return -1, err
	}

	return data, nil
}
