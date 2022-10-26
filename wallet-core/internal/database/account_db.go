package database

import (
	"database/sql"

	"github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"
)

type AccountDB struct {
	Db *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{
		Db: db,
	}
}

func (a AccountDB) FindById(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	stmt, err := a.Db.Prepare("SELECT a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at FROM accounts a INNER JOIN clients c ON a.client_id = c.id WHERE a.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.Id,
		&account.Client.Id,
		&account.Balance,
		&account.CreatedAt,
		&client.Id,
		&client.Name,
		&client.Email,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a AccountDB) Save(account *entity.Account) error {
	stmt, err := a.Db.Prepare("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.Id, account.Client.Id, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
