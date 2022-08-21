package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transfer struct {
	ID     string
	Amount float64
	Status string
	// Description  string
	// Store        string
	// CreditCardId string
	AccountIdFrom string
	AccountIdTo   string
	CreatedAt     time.Time
}

func NewTransfer() *Transfer {
	t := &Transfer{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	t.AccountIdFrom = accountFrom
	t.AccountIdTo = accountTo
	return t
}

func (t *Transfer) ProcessAndValidate(accountTo *Account, accountFrom *Account) {
	if t.Amount > accountFrom.Balance {
		t.Status = "rejected"
	} else {
		t.Status = "approved"
		accountTo.Balance = accountTo.Balance + t.Amount
	}
}
