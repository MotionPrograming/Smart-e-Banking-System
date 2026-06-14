package user

import (
	"net/http"
	"strings"

	"smart-e-banking/backend/util"
)

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimSpace(r.URL.Query().Get("email"))
	if email == "" {
		util.SendError(w, http.StatusBadRequest, "email query parameter is required")
		return
	}

	usr, err := h.svc.GetUserByEmail(email)
	if err != nil {
		util.SendError(w, http.StatusNotFound, "user not found")
		return
	}

	util.SendData(w, http.StatusOK, map[string]interface{}{
		"id":    usr.ID,
		"name":  usr.Name,
		"email": usr.Email,
		"phone": usr.Phone,
		"role":  usr.Role,
	})
}
