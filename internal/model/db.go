package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dsn string) error {
	dsn := "root:bjgkSfio_21751904fordemacia@tcp(localhost:3306)/idle_game?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil
}
