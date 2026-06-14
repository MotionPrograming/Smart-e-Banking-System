package account

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	_ = id // call service here later

	w.WriteHeader(http.StatusOK)
}
