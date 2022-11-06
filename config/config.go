package config

import (
	"database/sql"
	"fmt"

	"github.com/microsoft/go-mssqldb/azuread"
)

const (
	username string = "your username database"
	password string = "your password database"
	database string = "your name database"
)

var (
	dsn = fmt.Sprintf("sqlserver://%s:%s@your_ip:your_port?database=%s", username, password)
)

func MsSql() (*sql.DB, error) {
	db, err := sql.Open(azuread.DriverName, dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}
