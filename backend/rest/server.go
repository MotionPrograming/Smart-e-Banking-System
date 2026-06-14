package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"smart-e-banking/backend/config"
	account "smart-e-banking/backend/rest/handlers/account"
	transaction "smart-e-banking/backend/rest/handlers/transaction"
	user "smart-e-banking/backend/rest/handlers/user"
	"smart-e-banking/backend/rest/middleware"
)

type Server struct {
	cnf                *config.Config
	userHandler        *user.Handler
	transactionHandler *transaction.Handler
	accountHandler     *account.Handler
}

func NewServer(
	cnf *config.Config,
	userHandler *user.Handler,
	transactionHandler *transaction.Handler,
	accountHandler *account.Handler,
) *Server {
	return &Server{
		cnf:                cnf,
		userHandler:        userHandler,
		transactionHandler: transactionHandler,
		accountHandler:     accountHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()

	server.userHandler.RegisterRoutes(mux, manager)
	server.transactionHandler.RegisterRoutes(mux, manager)
	server.accountHandler.RegisterRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)

	addr := ":" + strconv.Itoa(server.cnf.HTTPPort)
	fmt.Printf("Server running on: %s\n", addr)

	srv := &http.Server{
		Addr:         addr,
		Handler:      wrappedMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	fmt.Println("Shutting down server...")

	srv.Shutdown(context.Background())
}
