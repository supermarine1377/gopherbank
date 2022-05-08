package infrastructure

import (
	"database/sql"
	"flag"
	"fmt"
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

type Transaction struct {
	tx *sql.Tx
}

func NewSqlHandler() *SqlHandler {
	host := getHost()
	dataSourceName := fmt.Sprintf("root:password@tcp(%s:3306)/gopherbank", host)
	comn, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return &SqlHandler{Comn: comn}
}

func (handler *SqlHandler) Excute(statement string, args ...interface{}) (db.Result, error) {
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

func (handler *SqlHandler) Query(statement string, args ...interface{}) (db.Row, error) {
	// rows, err := tx.Query(statement, args...)
	rows, err := handler.Comn.Query(statement, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return SqlRows{Rows: rows}, nil
}

func (sqlResult SqlResult) LastInsertId() (int64, error) {
	return sqlResult.LastInsertId()
}

func (sqlResult SqlResult) RowsAffected() (int64, error) {
	return sqlResult.RowsAffected()
}

func (sqlRows SqlRows) Scan(dest ...interface{}) error {
	return sqlRows.Rows.Scan(dest...)
}

func (sqlRows SqlRows) Next() bool {
	return sqlRows.Rows.Next()
}

func (sqlRows SqlRows) Close() error {
	return sqlRows.Rows.Close()
}

func (handler *SqlHandler) DoInTx(f func() (interface{}, error)) (interface{}, error) {
	tx, err := handler.Comn.Begin()
	defer func() {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	r, err := f()
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return r, nil
}

// docker-composeから起動した場合、"gopher-db"を返す
// go run . で起動した場合、"localhost"を返す (CI/CDのテスト用)
func getHost() string {
	p := flag.String("db_host", "localhost", "MySQL host name")
	flag.Parse()
	return *p
}
