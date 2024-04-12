package reports

import (
	"context"
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type ReportsStore struct {
	driver string
	path   string
}

func NewReportStore() *ReportsStore {
	repStore := ReportsStore{
		path:   "reports",
		driver: "sqlite",
	}

	if err := repStore.Ping(); err != nil {
		panic(err)
	}

	return &repStore
}

func (r *ReportsStore) Ping() error {
	db, err := sql.Open(r.driver, r.path)
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}

func (r *ReportsStore) GetQuery(ctx context.Context, id string) (string, error) {
	// log.Print(id)

	db, err := sql.Open(r.driver, r.path)
	if err != nil {
		return err.Error(), err
	}
	defer db.Close()

	conn, err := db.Conn(ctx)
	if err != nil {
		return err.Error(), err
	}

	var data string

	//id is an argument !!!
	// q := db.QueryRow()
	row := conn.QueryRowContext(ctx, "SELECT query FROM reports WHERE id=$1;", id)
	err = row.Scan(&data)
	if err != nil {
		return err.Error(), err
	}

	return data, nil

}
