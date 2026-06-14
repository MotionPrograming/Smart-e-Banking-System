package user

import (
	"net/http"
	"strconv"

	"smart-e-banking/backend/util"
)

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		util.SendError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	usr, err := h.svc.GetUserByID(id)
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
