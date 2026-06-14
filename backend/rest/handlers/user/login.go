package user

import (
	"encoding/json"
	"net/http"
	"time"

	"smart-e-banking/backend/util"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var req ReqLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usr, err := h.svc.Authenticate(req.Email, req.Password)
	if err != nil {

		util.SendError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	payload := util.Payload{
		Sub:         usr.ID,
		Name:        usr.Name,
		Email:       usr.Email,
		IsShopOwner: usr.Role == "shop_owner",
		IssuedAt:    time.Now().Unix(),
		ExpiresAt:   time.Now().Add(24 * time.Hour).Unix(),
	}

	token, err := util.CreateJWT(payload, h.cnf.JWTSecretKey)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	util.SendData(w, http.StatusOK, map[string]interface{}{
		"access_token": token,
		"user": map[string]interface{}{
			"id":            usr.ID,
			"name":          usr.Name,
			"email":         usr.Email,
			"is_shop_owner": usr.Role == "shop_owner",
		},
	})
}
