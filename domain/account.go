package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID string
	// Name            string
	AccountNumber string
	// ExpirationMonth int32
	// ExpirationYear  int32
	// CVV             int32
	Balance float64
	// Limit           float64
	CreatedAt time.Time
}

func NewAccount() *Account {
	c := &Account{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()
	return c
}
