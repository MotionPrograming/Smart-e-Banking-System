package account

import (
	"encoding/json"
	"net/http"
)

type reqUpdateAccount struct {
	AccountType string `json:"account_type"`
	Status      string `json:"status"`
}

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req reqUpdateAccount
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	acc, err := h.svc.GetAccountByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if req.AccountType != "" {
		acc.AccountType = req.AccountType
	}
	if req.Status != "" {
		acc.Status = req.Status
	}

	if err := h.svc.UpdateAccount(acc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}
