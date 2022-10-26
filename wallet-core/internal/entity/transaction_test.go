package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("client1", "Client@1.com")
	account1 := NewAccount(client1)
	account1.Credit(1000.0)

	client2, _ := NewClient("client2", "Client@2.com")
	account2 := NewAccount(client2)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 900.0, account1.Balance)
	assert.Equal(t, 1100.0, account2.Balance)

}

func TestCreateTransactionWithInvalidAmount(t *testing.T) {
	client1, _ := NewClient("client1", "Client@1.com")
	account1 := NewAccount(client1)
	account1.Credit(1000.0)

	client2, _ := NewClient("client2", "Client@2.com")
	account2 := NewAccount(client2)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, -100.0)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}

func TestCreateTransactionWithInsufficientFounds(t *testing.T) {
	client1, _ := NewClient("client1", "Client@1.com")
	account1 := NewAccount(client1)
	account1.Credit(1000.0)

	client2, _ := NewClient("client2", "Client@2.com")
	account2 := NewAccount(client2)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 1100.0)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient founds")
	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}
