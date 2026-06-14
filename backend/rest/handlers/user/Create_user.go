package user

import (
	"encoding/json"
	"net/http"
	"smart-e-banking/backend/domain"
	"smart-e-banking/backend/util"
	"time"
)

type ReqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateUser

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Request Data")
		return
	}
	usr, err := h.svc.CreateUser(domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
		Phone:        req.Phone,
		Role:         req.Role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	util.SendData(w, http.StatusCreated, usr)
}
