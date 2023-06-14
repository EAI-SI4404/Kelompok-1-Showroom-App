package routes

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/controllers"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupPaymentRoutes(app *fiber.App) {

	payment := app.Group("/payment").Use(middleware.CustomerAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	payment.Post("/create", controllers.CreatePayment)
	payment.Get("/method", controllers.GetAllPaymentMethods)

}
