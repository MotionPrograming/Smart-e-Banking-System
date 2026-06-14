package account

import (
	"net/http"
	"smart-e-banking/backend/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /accounts", manager.With(http.HandlerFunc(h.CreateAccount)))
	mux.Handle("GET /accounts/{id}", manager.With(http.HandlerFunc(h.GetAccountByID)))
	mux.Handle("PUT /accounts/{id}", manager.With(http.HandlerFunc(h.UpdateAccount)))
	mux.Handle("DELETE /accounts/{id}", manager.With(http.HandlerFunc(h.DeleteAccount)))
	mux.Handle("GET /accounts/{id}/users/{userID}", manager.With(http.HandlerFunc(h.GetAccountsByUserID)))

}
