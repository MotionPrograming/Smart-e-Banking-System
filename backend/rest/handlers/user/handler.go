package user

import (
	"smart-e-banking/backend/config"
	svc "smart-e-banking/backend/service/user"
)

type Handler struct {
	cnf *config.Config
	svc svc.Service
}

func NewHandler(cnf *config.Config, svc svc.Service) *Handler {
	return &Handler{cnf: cnf, svc: svc}
}
