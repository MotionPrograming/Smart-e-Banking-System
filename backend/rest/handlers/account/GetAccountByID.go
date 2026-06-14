package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetAccountByID(w http.ResponseWriter, r *http.Request) {
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

	account, err := h.svc.GetAccountByID(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
