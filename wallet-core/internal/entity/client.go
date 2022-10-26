package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("Name is required")
	}
	if c.Email == "" {
		return errors.New("Email is required")
	}
	return nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.Id != c.Id {
		return errors.New("Account does not belong to client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}
