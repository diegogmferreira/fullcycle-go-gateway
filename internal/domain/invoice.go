package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Status         Status
	Amount         float64
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpireMonth    int
	ExpireYear     int
	CardholderName string
}

func NewInvoice(accoundID string, ammount float64, description string, paymentType string, card *CreditCard) (*Invoice, error) {
	if ammount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accoundID,
		Amount:         ammount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()

	return nil
}

func (i *Invoice) UpdateStatus(status Status) error {
	if i.Status != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = status
	i.UpdatedAt = time.Now()

	return nil
}
