package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)

type WithdrawRequest struct {
	Amount decimal.Decimal `json:"amount"`
}

// assuming:
// type Handler struct {
//     svc Service
// }

func (h *Handler) Withdraw(w http.ResponseWriter, r *http.Request) {

	// 1. get account id from url
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
	var req WithdrawRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// 3. call service
	err = h.svc.Withdraw(accountID, req.Amount, " withdrawal from REST API")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4. success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "withdraw successful",
	})
}
