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
	mux.Handle("GET /users/{user_id}/accounts", manager.With(http.HandlerFunc(h.GetAccountsByUserID)))
	mux.Handle("POST /accounts/{id}/deposit", manager.With(http.HandlerFunc(h.Deposit)))
	mux.Handle("POST /accounts/{id}/withdraw", manager.With(http.HandlerFunc(h.Withdraw)))
	mux.Handle("POST /accounts/transfer", manager.With(http.HandlerFunc(h.Transfer)))
}
