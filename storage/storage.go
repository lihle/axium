package storage

import (
	"database/sql"
	//this driver is needed to run sql queries to mysql

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//Init connects to the database and returns an error upon failure
func Init() error {
	var err error
	db, err = conn()
	return err
}

func conn() (*sql.DB, error) {
	return sql.Open("mysql", "lihle:lihle@(localhost:3306)/new_datalicious?charset=utf8")
}

func valueOrWildcard(in string) string {
	if in == "" {
		return "%"
	}
	return in
}
