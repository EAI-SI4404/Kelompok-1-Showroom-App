package controllers

import (
	"net/http"

	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *BalanceController) TopUpBalance(ctx *fiber.Ctx) error {
	customerToken := ctx.Locals("customer").(domain.CustomerToken)
	req := dto.TopUpBalanceRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	customer, err := c.balanceService.FindCustomerByToken(customerToken.Token)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find customer",
		})
	}

	// get customer balance
	customerBalance, err := c.balanceService.GetBalance(customer.ID)

	if err != nil {
		// create new balance
		customerBalance = &domain.CutstomerBalance{
			CustomerID: customer.ID,
			Balance:    req.Amount,
		}

		customerBalance, err = c.balanceService.TopUpBalance(customerBalance)

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to top up balance",
			})
		}
	} else {
		// update balance
		customerBalance.Balance += req.Amount

		customerBalance, err = c.balanceService.UpdateBalance(customerBalance)

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to top up balance",
			})
		}
	}

	// response
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success top up balance",
		"data": fiber.Map{
			"balance": customerBalance.Balance,
		},
	})

}

func (c *BalanceController) GetBalance(ctx *fiber.Ctx) error {
	// get customer token
	customerToken := ctx.Locals("customer").(domain.CustomerToken)

	// find customer by token
	customer, err := c.balanceService.FindCustomerByToken(customerToken.Token)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to find customer",
		})
	}

	// get customer balance
	customerBalance, err := c.balanceService.GetBalance(customer.ID)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get balance",
		})
	}

	// response
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get balance",
		"data": fiber.Map{
			"balance": customerBalance.Balance,
		},
	})
}
