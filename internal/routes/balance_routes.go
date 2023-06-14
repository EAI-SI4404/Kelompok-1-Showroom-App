package routes

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/controllers"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/middleware"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupBalanceRoutes(router fiber.Router, balanceService services.BalanceService) {
	balanceController := controllers.NewBalanceHandler(balanceService)

	balance := router.Group("/balance").Use(middleware.CustomerAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))

	balance.Post("/topup", balanceController.TopUpBalance)
	balance.Get("", balanceController.GetBalance)

}
