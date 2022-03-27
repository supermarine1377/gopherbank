package infrastructure

import (
	"database/sql"
	"log"
	"supermarine1377/interface/db"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Comn *sql.DB
}

type SqlResult struct {
	Result *sql.Result
}

type SqlRows struct {
	Rows *sql.Rows
}

func NewSqlHandler() *SqlHandler {
	comn, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/gopherbank")
	if err != nil {
		panic(err)
	}
	return &SqlHandler{Comn: comn}
}

func (handler SqlHandler) Excute(statement string, args ...interface{}) (db.Result, error) {
	tx, err := handler.Comn.Begin()
	defer tx.Rollback()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err := tx.Exec(statement, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (handler SqlHandler) Query(statement string, args ...interface{}) (db.Row, error) {
	tx, err := handler.Comn.Begin()
	defer tx.Rollback()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := tx.Query(statement, args)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		log.Println(err)
		return nil, err
	}
	return SqlRows{Rows: rows}, nil
}

func (sqlRows SqlRows) LastInsertId() (int64, error) {
	return sqlRows.LastInsertId()
}

func (sqlRows SqlRows) RowsAffected() (int64, error) {
	return sqlRows.RowsAffected()
}

func (sqlRows SqlRows) Scan(dest ...interface{}) error {
	return sqlRows.Rows.Scan(dest)
}

func (sqlRows SqlRows) Next() bool {
	return sqlRows.Rows.Next()
}

func (sqlRows SqlRows) Close() error {
	return sqlRows.Rows.Close()
}
