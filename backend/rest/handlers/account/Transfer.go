package account

import (
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
)

type reqTransfer struct {
	FromAccountID int64           `json:"from_account_id"`
	ToAccountID   int64           `json:"to_account_id"`
	Amount        decimal.Decimal `json:"amount"`
}

func (h *Handler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req reqTransfer
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.svc.Transfer(req.FromAccountID, req.ToAccountID, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "transfer successful"})
}
