package main

import (
	"fmt"
	"supermarine1377/infrastructure"
)

func main() {
	fmt.Println("starting server...")
	var (
		sql = infrastructure.NewSqlHandler()
		db  = sql.Comn
	)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5)
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
	infrastructure.Route(sql)
}
