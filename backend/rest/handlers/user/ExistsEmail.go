package user

import (
	"net/http"
	"smart-e-banking/backend/util"
)

func (h *Handler) EmailExists(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")

	exists, err := h.svc.EmailExists(email)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	util.SendData(w, http.StatusOK, map[string]interface{}{
		"exists": exists,
	})
}
