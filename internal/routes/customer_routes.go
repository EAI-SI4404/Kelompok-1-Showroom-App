package routes

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/controllers"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupCustomerRoutes(router fiber.Router, customerService services.CustomerService) {
	authController := controllers.NewUserHandler(customerService)

	router.Post("/register", authController.Register)

}
