package controllers

import (
	"net/http"

	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *CustomerController) Register(ctx *fiber.Ctx) error {
	req := dto.RegisterRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	role, err := c.customerService.FindRoleByName("customer")

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find role",
		})

	}

	customer := &domain.Customer{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
		Address:  req.Address,
		City:     req.City,
		Province: req.Province,
		ZipCode:  req.ZipCode,
		Role:     role.Name,
	}

	customer, err = c.customerService.Register(customer)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return Response
	response := dto.RegisterResponse{
		ID:       customer.ID,
		Name:     customer.Name,
		Email:    customer.Email,
		Phone:    customer.Phone,
		Address:  customer.Address,
		City:     customer.City,
		Province: customer.Province,
		ZipCode:  customer.ZipCode,
		Role:     customer.Role,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Register success",
		"data":    response,
	})

}
