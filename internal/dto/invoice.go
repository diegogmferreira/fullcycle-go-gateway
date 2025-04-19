package dto

import (
	"time"

	"github.com/diegogmferreira/fullcycle-go-gateway/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceRequest struct {
	APIKey         string
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpireMonth    int     `json:"expire_month"`
	ExpireYear     int     `json:"expire_year"`
	CardholderName string  `json:"cardholder_name"`
}

type InvoiceResponse struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(request CreateInvoiceRequest, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:         request.CardNumber,
		CVV:            request.CVV,
		ExpireMonth:    request.ExpireMonth,
		ExpireYear:     request.ExpireYear,
		CardholderName: request.CardholderName,
	}

	return domain.NewInvoice(
		accountID,
		request.Amount,
		request.Description,
		request.PaymentType,
		&card,
	)
}

func ToInvoiceResponse(invoice *domain.Invoice) *InvoiceResponse {
	return &InvoiceResponse{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
