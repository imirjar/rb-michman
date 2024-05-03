package reports

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/imirjar/Michman/internal/diver/models"
	_ "modernc.org/sqlite"
)

type ReportsStore struct {
	db *sql.DB
}

func NewReportStore() *ReportsStore {
	db, err := sql.Open("sqlite", "reports")
	if err != nil {
		panic(err)
	}

	repStore := ReportsStore{
		db: db,
	}

	if err := repStore.Ping(); err != nil {
		panic(err)
	}

	return &repStore
}

func (r *ReportsStore) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	return r.db.PingContext(ctx)
}

func (r *ReportsStore) GetQuery(ctx context.Context, id string) (string, error) {
	// log.Print(id)

	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err.Error(), err
	}
	defer conn.Close()

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

func (r *ReportsStore) GetAllReports(ctx context.Context) (string, error) {
	conn, err := r.db.Conn(ctx)
	if err != nil {
		return err.Error(), err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT * FROM reports;")
	if err != nil {
		return err.Error(), err
	}

	var reports []models.Report
	for rows.Next() {

		var rep models.Report
		err = rows.Scan(&rep.Id, &rep.Name, &rep.Query)
		if err != nil {
			return err.Error(), err
		}
		reports = append(reports, rep)
	}
	result, err := json.Marshal(&reports)
	if err != nil {
		return err.Error(), err
	}

	return string(result), nil
}
