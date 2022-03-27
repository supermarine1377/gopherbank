package transaction

import (
	"supermarine1377/domain"
	"supermarine1377/interface/db"
)

type TransactionRepository struct {
	sqlHandler db.SqlHandler
}

func (repo TransactionRepository) Store(t domain.Transaction) {

}
