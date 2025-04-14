package service

import (
	"github.com/diegogmferreira/fullcycle-go-gateway/internal/domain"
	"github.com/diegogmferreira/fullcycle-go-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (s *AccountService) CreateAccount(request dto.CreateAccountRequest) (*dto.AccountResponse, error) {
	account := dto.ToAccount(request)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}

	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	output := dto.ToAccountResponse(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.ToAccountResponse(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.ToAccountResponse(account)
	return &output, nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {

	}

	output := dto.ToAccountResponse(account)
	return &output, nil
}
