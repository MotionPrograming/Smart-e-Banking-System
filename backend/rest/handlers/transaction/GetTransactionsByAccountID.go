package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetTransactionsByAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing account ID", http.StatusBadRequest)
		return
	}

	accountID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || accountID <= 0 {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	transactions, err := h.svc.GetTransactionsByAccountID(accountID)
	if err != nil {
		http.Error(w, "Failed to retrieve transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
