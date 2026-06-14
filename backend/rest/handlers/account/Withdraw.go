package account

import (
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
)

type reqWithdraw struct {
	Amount    decimal.Decimal `json:"amount"`
	Reference string          `json:"reference"`
}

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {
	accountID, err := parseID(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req reqWithdraw
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.svc.Withdraw(accountID, req.Amount, req.Reference); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "withdrawal successful"})
}
