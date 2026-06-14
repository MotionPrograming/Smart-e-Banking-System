package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	fmt.Println("Server running on:", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
