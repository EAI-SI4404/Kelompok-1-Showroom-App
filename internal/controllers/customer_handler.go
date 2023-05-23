package controllers

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/services"
)

type CustomerController struct {
	customerService services.CustomerService
}

func NewUserHandler(customerService services.CustomerService) *CustomerController {
	return &CustomerController{customerService}
}
