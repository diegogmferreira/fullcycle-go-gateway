package service

import (
	"github.com/diegogmferreira/fullcycle-go-gateway/internal/domain"
	"github.com/diegogmferreira/fullcycle-go-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountService:    accountService,
	}
}

func (s *InvoiceService) Create(request dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error) {
	accountOutput, err := s.accountService.FindByAPIKey(request.APIKey)

	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(request, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(request.APIKey, invoice.Amount)

		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.ToInvoiceResponse(invoice), nil
}

func (s *InvoiceService) GetByID(id, apiKey string) (*dto.InvoiceResponse, error) {
	invoice, err := s.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	accountOutput, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != accountOutput.ID {
		return nil, domain.ErrUnauthorizedAccess
	}

	return dto.ToInvoiceResponse(invoice), nil
}

func (s *InvoiceService) ListByAccount(accoundID string) ([]*dto.InvoiceResponse, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accoundID)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceResponse, 0)
	for i, invoice := range invoices {
		output[i] = dto.ToInvoiceResponse(invoice)
	}

	return output, nil
}

func (s *InvoiceService) ListByAccountAPIKey(apiKey string) ([]*dto.InvoiceResponse, error) {
	accountOutput, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccount(accountOutput.ID)
}
