package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diegogmferreira/fullcycle-go-gateway/internal/domain"
	"github.com/diegogmferreira/fullcycle-go-gateway/internal/dto"
	"github.com/diegogmferreira/fullcycle-go-gateway/internal/service"
	"github.com/go-chi/chi/v5"
)

type InvoiceHandler struct {
	service *service.InvoiceService
}

func NewInvoiceHandler(service *service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		service: service,
	}
}

// Request authenticated by X-API-KEY
// Endpoint: /invoice
// Method: POST
func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateInvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.APIKey = r.Header.Get("X-API-KEY")

	response, err := h.service.Create(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Request authenticated by X-API-KEY
// Endpoint: /invoice/{id}
// Method: GET
func (h *InvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	apiKey := r.Header.Get("X-API-KEY")
	if apiKey == "" {
		http.Error(w, "X-API-KEY is required", http.StatusBadRequest)
		return
	}

	response, err := h.service.GetByID(id, apiKey)
	if err != nil {
		switch err {
		case domain.ErrInvoiceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case domain.ErrUnauthorizedAccess:
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Request authenticated by X-API-KEY
// Endpoint: /invoice/list
// Method: GET
func (h *InvoiceHandler) ListByAccount(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("X-API-KEY")
	if apiKey == "" {
		http.Error(w, "X-API-KEY is required", http.StatusBadRequest)
		return
	}

	response, err := h.service.ListByAccountAPIKey(apiKey)
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
