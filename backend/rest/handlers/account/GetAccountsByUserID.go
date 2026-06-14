package account

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAccountsByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := parseID(r, "user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accounts, err := h.svc.GetAccountsByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
