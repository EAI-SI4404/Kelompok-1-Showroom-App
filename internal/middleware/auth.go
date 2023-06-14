package middleware

import (
	"log"

	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/config"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type AuthConfig struct {
	Filter       func(*fiber.Ctx) error
	Unauthorized fiber.Handler
}

func CustomerAuthentication(c AuthConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header := ctx.GetReqHeaders()
		// check if authorization header is exist

		log.Println(header)

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		// check user token is valid
		customerToken := domain.CustomerToken{}

		err := config.DB.Where("token = ?", header["Authorization"]).First(&customerToken).Error

		if err != nil {
			return c.Unauthorized(ctx)
		}

		if customerToken.Role != "customer" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("customer", customerToken)

		return ctx.Next()
	}
}

func AdminAuthentication(c AuthConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		// check admin token is valid
		var adminToken domain.CustomerToken

		err := config.DB.Where("token = ?", header["Authorization"]).First(&adminToken).Error

		if err != nil {
			return c.Unauthorized(ctx)
		}

		if adminToken.Role != "admin" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("admin", adminToken.Token)

		return ctx.Next()

	}
}
