package user // প্যাকেজ নাম পরিবর্তন করে 'user' দিন

import (
	"net/http"
	"smart-e-banking/backend/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Go 1.22+ এর নতুন রাউটিং সিনট্যাক্স ব্যবহার করছেন আপনি
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
	mux.Handle("GET /users/email", manager.With(http.HandlerFunc(h.GetUserByEmail)))
	mux.Handle("GET /users/id", manager.With(http.HandlerFunc(h.GetUserByID)))
	mux.Handle("GET /users/exists", manager.With(http.HandlerFunc(h.EmailExists)))
}
