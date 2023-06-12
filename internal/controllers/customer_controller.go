package controllers

import (
	"net/http"
	"time"

	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/dto"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/utils"
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

func (c *CustomerController) Login(ctx *fiber.Ctx) error {
	req := dto.LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	customer, err := c.customerService.Login(req.Email, req.Password)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	token, err := utils.CreateUserToken(customer)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create token",
		})
	}

	// Set cookie
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	// return Response
	response := dto.LoginResponse{
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

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login success",
		"data":    response,
		"token":   token,
	})

}

func (c *CustomerController) GetCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	customer, err := c.customerService.GetCustomer(id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.GetCustomerResponse{
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

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get customer",
		"data":    response,
	})
}

func (c *CustomerController) GetAllCustomer(ctx *fiber.Ctx) error {
	customers, err := c.customerService.GetAllCustomer()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var response []dto.GetCustomerResponse

	for _, customer := range customers {
		response = append(response, dto.GetCustomerResponse{
			ID:       customer.ID,
			Name:     customer.Name,
			Email:    customer.Email,
			Phone:    customer.Phone,
			Address:  customer.Address,
			City:     customer.City,
			Province: customer.Province,
			ZipCode:  customer.ZipCode,
			Role:     customer.Role,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get all customer",
		"data":    response,
	})
}
