package transaction

import (
	"encoding/json"
	"net/http"
	"smart-e-banking/backend/domain"

	"github.com/shopspring/decimal"
)

type reqCreateTransaction struct {
	FromAccountID *int64          `json:"from_account_id"`
	ToAccountID   *int64          `json:"to_account_id"`
	Amount        decimal.Decimal `json:"amount"`
	Currency      string          `json:"currency"`
	Type          string          `json:"type"`
	Reference     string          `json:"reference"`
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req reqCreateTransaction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	validTypes := map[string]bool{"deposit": true, "withdraw": true, "transfer": true}
	if !validTypes[req.Type] {
		http.Error(w, "invalid transaction type", http.StatusBadRequest)
		return
	}

	createdTx, err := h.svc.CreateTransaction(domain.Transaction{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
		Type:          req.Type,
		Reference:     req.Reference,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTx)
}
