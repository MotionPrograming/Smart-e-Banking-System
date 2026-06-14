package user

import (
	"net/http"
	"smart-e-banking/backend/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
	mux.Handle("GET /users/{id}", manager.With(http.HandlerFunc(h.GetUserByID)))
	mux.Handle("GET /users/by-email", manager.With(http.HandlerFunc(h.GetUserByEmail))) // ?email=
	mux.Handle("GET /users/exists", manager.With(http.HandlerFunc(h.EmailExists)))      // ?email=
}
