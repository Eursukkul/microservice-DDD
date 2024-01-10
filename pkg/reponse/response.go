package reponse

import "github.com/gofiber/fiber/v2"

type MsgResponse struct {
	Message string `json:"message"`
}

func ErrResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(&MsgResponse{Message: message})
}

func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(data)
}
