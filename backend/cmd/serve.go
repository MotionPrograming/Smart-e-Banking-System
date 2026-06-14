package cmd

import (
	"fmt"
	"log"
	"os"

	"smart-e-banking/backend/config"
	"smart-e-banking/backend/infra/db"
	server "smart-e-banking/backend/rest"

	accountRepoPkg "smart-e-banking/backend/repository/account"
	transactionRepoPkg "smart-e-banking/backend/repository/transaction"
	userRepoPkg "smart-e-banking/backend/repository/user"

	accountService "smart-e-banking/backend/service/account"
	transactionService "smart-e-banking/backend/service/transaction"
	userService "smart-e-banking/backend/service/user"

	accountHandler "smart-e-banking/backend/rest/handlers/account"
	transactionHandler "smart-e-banking/backend/rest/handlers/transaction"
	userHandler "smart-e-banking/backend/rest/handlers/user"
)

func Serv() {
	// Load config
	cnf := config.GetConfig()

	fmt.Printf("%+v\n", cnf.DB)

	// Database connection
	conn, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Run migrations
	err = db.MigrateDB(conn, "./backend/migration")
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// ======================
	// Repositories
	// ======================

	usrRepo := userRepoPkg.NewUserRepository(conn)
	accRepo := accountRepoPkg.NewAccountRepository(conn)
	transRepo := transactionRepoPkg.NewTransactionRepository(conn)

	usrSvc := userService.NewService(usrRepo)

	// IMPORTANT:
	// Transaction service needs BOTH repositories

	accSvc := accountService.NewService(
		conn, accRepo, transRepo,
	)

	transSvc := transactionService.NewService(
		conn, transRepo, accRepo,
	)

	userH := userHandler.NewHandler(cnf, usrSvc)

	accH := accountHandler.NewHandler(cnf, accSvc)

	transH := transactionHandler.NewHandler(cnf, transSvc)

	srv := server.NewServer(
		cnf,
		userH,
		transH,
		accH,
	)

	srv.Start()
}
