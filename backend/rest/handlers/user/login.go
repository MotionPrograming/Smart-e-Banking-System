package user

import (
	"encoding/json"
	"net/http"
	"time"

	"smart-e-banking/backend/util"

	"github.com/golang-jwt/jwt/v5"
)

type reqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req reqLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	usr, err := h.svc.Authenticate(req.Email, req.Password)
	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	payload := util.Payload{
		Sub:   usr.ID,
		Name:  usr.Name,
		Email: usr.Email,
		Role:  usr.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token, err := util.CreateJWT(payload, h.cnf.JWTSecretKey)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	util.SendData(w, http.StatusOK, map[string]interface{}{
		"access_token": token,
		"user": map[string]interface{}{
			"id":    usr.ID,
			"name":  usr.Name,
			"email": usr.Email,
			"role":  usr.Role,
		},
	})
}
