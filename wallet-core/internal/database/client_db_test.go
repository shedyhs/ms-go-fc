package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDb *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDb = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		Name:  "Jhon Doe",
		Email: "j@j.com",
	}
	err := s.clientDb.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("name", "email@email.com")
	s.clientDb.Save(client)
	clientDB, err := s.clientDb.Get(client.Id)
	s.Nil(err)
	s.Equal(client.Id, clientDB.Id)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)

}
