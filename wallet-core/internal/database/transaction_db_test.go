package database

import (
	"database/sql"
	"testing"

	"github.com.br/shedyhs/ms-fullcycle/wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDbTestSuite struct {
	suite.Suite
	db *sql.DB

	clientFrom  *entity.Client
	accountFrom *entity.Account

	clientTo  *entity.Client
	accountTo *entity.Account

	transactionDB *TransactionDB
}

func (s *TransactionDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")
	// Creating clients
	clientFrom, err := entity.NewClient("Jhon Doe", "j@j.com")
	s.Nil(err)
	s.clientFrom = clientFrom

	clientTo, err := entity.NewClient("Doe Jhon", "a@b.com")
	s.Nil(err)
	s.clientTo = clientTo
	// Creating accounts
	accountFrom := entity.NewAccount(clientFrom)
	s.NotNil(accountFrom)
	accountFrom.Balance = 1000
	s.accountFrom = accountFrom

	accountTo := entity.NewAccount(clientTo)
	s.NotNil(accountTo)
	accountTo.Balance = 1000
	s.accountTo = accountTo

	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDbTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDbTestSuite))
}

func (s *TransactionDbTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
