package handlers

import (
	"log"
	"net/http"
	"sample/db"
	"sample/models"

	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	_, err := db.DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", newUser.Username, newUser.Email)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{"message": "User added successfully", "user": newUser})
}

// Pays handles the POST request to insert data into the trytable.
// @Summary Check if the user is online or offline based on input parameters.
// @Description Inserts user data into the trytable and provides a response message.
// @ID Post-transfer
// @Accept json
// @Produce json
// @Param request body models.TransferRequest true "JSON request body"
// @Success 200 {object} models.TransferRequest
// @Router /transfer [post]

func TransCredit(c *fiber.Ctx) error {
	transaction := &models.TransferRequest{}

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "JSON parsing error",
			"details": err.Error(),
		})
	}

	iGate := "http://127.0.0.1:1432/api/v1/ips/fdsaps"

	req, err := http.NewRequest(http.MethodPost, iGate, nil)

	if err != nil {
		log.Printf("Error creating request: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Merchant-ID", "QVBJMDAwMDU=")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Header": req.Header,
			"data":   transaction,
		})
	}
	return c.JSON(transaction)
}
