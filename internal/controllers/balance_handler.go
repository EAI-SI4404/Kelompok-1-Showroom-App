package controllers

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/services"
)

type BalanceController struct {
	balanceService services.BalanceService
}

func NewBalanceHandler(balanceService services.BalanceService) *BalanceController {
	return &BalanceController{balanceService}
}
