package main

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	// start application
	app.StartApplication()
}
