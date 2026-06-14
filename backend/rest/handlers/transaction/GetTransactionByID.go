package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing transaction ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	tx, err := h.svc.GetTransactionByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if tx == nil {
		http.Error(w, "Transaction not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tx); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
