package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Jhon Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Jhon Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	err := client.Update("Jhon Doe Updated", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "Jhon Doe Updated", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	err := client.Update("", "j@j.com")
	assert.Error(t, err, "Name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccountToAnotherClient(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	client2, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	err := client2.AddAccount(account)
	assert.Error(t, err, "Account does not belong to client")
	assert.Equal(t, 0, len(client2.Accounts))
}
