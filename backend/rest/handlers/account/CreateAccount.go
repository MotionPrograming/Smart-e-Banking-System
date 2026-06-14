package account

import (
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
)

type reqCreateAccount struct {
	UserID         int64           `json:"user_id"`
	AccountType    string          `json:"account_type"`
	InitialBalance decimal.Decimal `json:"initial_balance"`
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req reqCreateAccount
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	acc, err := h.svc.CreateAccount(req.UserID, req.AccountType, req.InitialBalance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}
