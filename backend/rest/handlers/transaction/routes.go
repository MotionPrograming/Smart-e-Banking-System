package transaction

import (
	"net/http"
	"smart-e-banking/backend/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /transactions", manager.With(http.HandlerFunc(h.CreateTransaction)))
	mux.Handle("GET /transactions/{id}", manager.With(http.HandlerFunc(h.GetTransactionByID)))
	mux.Handle("GET /accounts/{id}/transactions", manager.With(http.HandlerFunc(h.GetTransactionsByAccountID)))
}
