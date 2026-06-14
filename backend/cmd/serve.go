package cmd

import (
	"fmt"
	"log"

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

	conn, err := db.NewConnection(cnf.DB)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	if err := db.MigrateDB(conn, "./backend/migration"); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	usrRepo := userRepoPkg.NewUserRepository(conn)
	accRepo := accountRepoPkg.NewAccountRepository(conn)
	transRepo := transactionRepoPkg.NewTransactionRepository(conn)

	usrSvc := userService.NewService(usrRepo)

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
