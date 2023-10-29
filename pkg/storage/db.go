package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func InitDb() error {
	d, err := sqlx.Connect("postgres", "user=dev password=abcde dbname=uber sslmode=disable")
	if err != nil {
		return err
	}
	Db = d
	Db.MustExec(schema)

	return nil
}
