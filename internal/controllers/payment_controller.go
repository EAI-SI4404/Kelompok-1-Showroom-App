package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// PaymentController is a struct that define the payment controller
// this controller is used to handle payment related request from external API

type PaymentType struct {
	AdminFee        int    `json:"admin_fee"`
	PaymentType     string `json:"payment_type"`
	PaymentTypeCode string `json:"payment_type_code"`
	PaymentTypeID   int    `json:"payment_type_id"`
}

type PaymentResponse struct {
	PaymentTypes []PaymentType `json:"paymentTypes"`
}

type CreatePaymentResponse struct {
	Message      string `json:"message"`
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	Timestamp    string `json:"timestamp"`
	CustomerName string `json:"customer_name"`
}

func GetAllPaymentMethods(ctx *fiber.Ctx) error {
	url := os.Getenv("PAYMENT_API") + "/getallpaymenttype"

	// Send a GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to call the API:", err)
		return err
	}
	defer resp.Body.Close()

	// Check if the API returned a successful response
	if resp.StatusCode != http.StatusOK {
		log.Println("API returned an error:", resp.Status)
		return fmt.Errorf("API returned an error: %s", resp.Status)
	}

	// Decode the response body into a PaymentResponse struct
	var paymentResponse PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&paymentResponse)
	if err != nil {
		log.Println("Failed to decode API response:", err)
		return err
	}

	// Return the payment types
	return ctx.JSON(paymentResponse)
}

func CreatePayment(ctx *fiber.Ctx) error {
	customerToken := ctx.Locals("customer").(domain.CustomerToken)
	url := os.Getenv("PAYMENT_API") + "/createpayment"

	// Send a POST request to the API
	req := dto.PaymentRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Convert the modified request to JSON
	reqJSON, err := json.Marshal(req)
	if err != nil {
		log.Println("Failed to marshal request:", err)
		return err
	}

	// Send the modified request to the API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		log.Println("Failed to call the API:", err)
		return err
	}

	defer resp.Body.Close()

	// Check if the API returned a successful response
	if resp.StatusCode != http.StatusOK {
		log.Println("API returned an error:", resp.Status)
		return fmt.Errorf("API returned an error: %s", resp.Status)
	}

	// Decode the response body into a CreatePaymentResponse struct
	var createPaymentResponse CreatePaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&createPaymentResponse)
	if err != nil {
		log.Printf("Error decoding CreatePaymentResponse")
		return err
	}

	// Assign the customer name to the CreatePaymentResponse struct
	createPaymentResponse.CustomerName = customerToken.Customer.Name

	return ctx.JSON(createPaymentResponse)
}

