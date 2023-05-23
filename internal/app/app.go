package app

import (
	"os"

	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/config"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/repositories"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/routes"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/services"
	"github.com/gofiber/fiber/v2"
)

func StartApplication() {
	app := fiber.New()

	// initialize database
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	// initialize repository
	customerRepository := repositories.NewCustomerRepository(db)

	// initialize service
	customerService := services.NewCustomerService(customerRepository)

	// initialize endpoint
	endpoint := app.Group("/api")

	// initialize routes
	routes.SetupCustomerRoutes(endpoint, customerService)

	// start application
	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
