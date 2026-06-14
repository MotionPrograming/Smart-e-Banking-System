package transaction

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid transaction id", http.StatusBadRequest)
		return
	}

	tx, err := h.svc.GetTransactionByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}
