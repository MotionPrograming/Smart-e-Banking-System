package transaction

import (
	"encoding/json"
	"net/http"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

type ReqCreateTransaction struct {
	FromAccountID int64   `json:"from_account_id"`
	ToAccountID   int64   `json:"to_account_id"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"` // e.g., "deposit", "withdrawal", "transfer"
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateTransaction

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// validate amount
	if req.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}

	// validate type
	validTypes := map[string]bool{
		"deposit":  true,
		"withdraw": true,
		"transfer": true,
	}

	if !validTypes[req.Type] {
		http.Error(w, "Invalid transaction type", http.StatusBadRequest)
		return
	}

	tx := domain.Transaction{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        decimal.NewFromFloat(req.Amount),
		Type:          req.Type,
	}

	createdTx, err := h.svc.CreateTransaction(tx)
	if err != nil {
		http.Error(w, "Failed to create transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTx)
}
