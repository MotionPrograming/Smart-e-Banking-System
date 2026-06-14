package user

import (
	"net/http"
	"smart-e-banking/backend/util"
	"strconv"
)

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.svc.GetUserByID(id)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if user == nil {
		util.SendError(w, http.StatusNotFound, "user not found")
		return
	}

	// safe response (no sensitive data leak)
	util.SendData(w, http.StatusOK, map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
		"role":  user.Role,
	})
}
