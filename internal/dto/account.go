package dto

import (
	"time"

	"github.com/diegogmferreira/fullcycle-go-gateway/internal/domain"
)

type CreateAccountRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	APIKey    string    `json:"api_key",omitempty`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccount(request CreateAccountRequest) *domain.Account {
	return domain.NewAccount(request.Name, request.Email)
}

func ToAccountResponse(account *domain.Account) AccountResponse {
	return AccountResponse{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		APIKey:    account.APIKey,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
