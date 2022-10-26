package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.Id, account.Client.Id)
}

func TestCreateAccountWithNilAccount(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	assert.Equal(t, float64(0), account.Balance)
	account.Credit(10)
	assert.Equal(t, float64(10), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	assert.Equal(t, float64(0), account.Balance)
	account.Credit(10)
	assert.Equal(t, float64(10), account.Balance)
	account.Debit(5)
	assert.Equal(t, float64(5), account.Balance)
}
