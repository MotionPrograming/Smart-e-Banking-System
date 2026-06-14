package user

import (
	"net/http"
	"strings"

	"smart-e-banking/backend/util"
)

func (h *Handler) EmailExists(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.URL.Query().Get("email"))
	if email == "" {
		util.SendError(w, http.StatusBadRequest, "email query parameter is required")
		return
	}

	exists, err := h.svc.EmailExists(email)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	util.SendData(w, http.StatusOK, map[string]bool{"exists": exists})
}
