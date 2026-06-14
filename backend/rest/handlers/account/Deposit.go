package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

type DepositRequest struct {
	Amount decimal.Decimal `json:"amount"`
}

func (h *Handler) Deposit(w http.ResponseWriter, r *http.Request) {

	// 1. get account id from URL
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing account id", http.StatusBadRequest)
		return
	}

	accountID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}

	// 2. parse request body
	var req DepositRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// 3. call service layer
	err = h.svc.Deposit(accountID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4. success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "deposit successful",
	})
}
