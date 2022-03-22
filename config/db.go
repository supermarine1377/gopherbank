package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	sql, err := sql.Open("mysql", "root@tcp(gopher-db)/gopherbank")
	if err != nil {
		panic(err)
	}
	return sql
}
