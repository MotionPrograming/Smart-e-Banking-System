package user

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"smart-e-banking/backend/domain"
	"smart-e-banking/backend/util"

	"golang.org/x/crypto/bcrypt"
)

type reqCreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req reqCreateUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	switch {
	case req.Name == "":
		util.SendError(w, http.StatusBadRequest, "name is required")
		return
	case !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, "."):
		util.SendError(w, http.StatusBadRequest, "a valid email address is required")
		return
	case len(req.Password) < 8:
		util.SendError(w, http.StatusBadRequest, "password must be at least 8 characters")
		return
	case len(req.Password) > 72:
		// bcrypt silently truncates at 72 bytes — reject early to avoid silent data loss.
		util.SendError(w, http.StatusBadRequest, "password must be at most 72 characters")
		return
	}

	exists, err := h.svc.EmailExists(req.Email)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if exists {
		util.SendError(w, http.StatusConflict, "email is already registered")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	role := req.Role
	if role == "" {
		role = "user"
	}

	usr, err := h.svc.CreateUser(domain.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Phone:        req.Phone,
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	util.SendData(w, http.StatusCreated, map[string]interface{}{
		"id":    usr.ID,
		"name":  usr.Name,
		"email": usr.Email,
		"phone": usr.Phone,
		"role":  usr.Role,
	})
}
