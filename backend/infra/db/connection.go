package db

import (
	"fmt"
	"smart-e-banking/backend/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetConnectionString(cnf *config.DBConfig) string {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
	)

	if cnf.EnableSSLMode {
		conn += "&tls=true"
	} else {
		conn += "&tls=false"
	}

	return conn
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", GetConnectionString(cnf))
	if err != nil {
		return nil, err
	}

	return db, nil
}
