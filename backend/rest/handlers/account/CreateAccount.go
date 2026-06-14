package account

import (
	"encoding/json"
	"net/http"

	"github.com/shopspring/decimal"
)

type ReqCreateAccount struct {
	UserID         int64           `json:"user_id"`
	AccountType    string          `json:"account_type"`
	InitialBalance decimal.Decimal `json:"initial_balance"`
}

type CreateAccountResponse struct {
	ID            int64           `json:"id"`
	UserID        int64           `json:"user_id"`
	AccountNumber string          `json:"account_number"`
	AccountType   string          `json:"account_type"`
	Balance       decimal.Decimal `json:"balance"`
	Status        string          `json:"status"`
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateAccount

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	account, err := h.svc.CreateAccount(
		req.UserID,
		req.AccountType,
		req.InitialBalance,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := CreateAccountResponse{
		ID:            account.ID,
		UserID:        account.UserID,
		AccountNumber: account.AccountNumber,
		AccountType:   account.AccountType,
		Balance:       account.Balance,
		Status:        account.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
