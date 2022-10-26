package database

import (
	"database/sql"

	"github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"
)

type ClientDB struct {
	Db *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{
		Db: db,
	}
}

func (c *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.Db.Prepare("SELECT id, name, email, created_at FROM clients WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	if err := row.Scan(&client.Id, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *ClientDB) Save(client *entity.Client) error {
	stmt, err := c.Db.Prepare("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.Id, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
