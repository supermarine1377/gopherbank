package db

type SqlHandler interface {
	Excute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

func DoInTx() {

}
