package user

import (
	"net/http"
	"smart-e-banking/backend/util"
)

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")

	usr, err := h.svc.GetUserByEmail(email)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	util.SendData(w, http.StatusOK, usr)
}
