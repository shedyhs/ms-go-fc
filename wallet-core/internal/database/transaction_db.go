package database

import (
	"database/sql"

	"github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"
)

type TransactionDB struct {
	Db *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		Db: db,
	}
}

func (t *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := t.Db.Prepare("INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		transaction.Id,
		transaction.AccountFrom.Id,
		transaction.AccountTo.Id,
		transaction.Amount,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
