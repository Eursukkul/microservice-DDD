package server

import (
	"github.com/gofiber/fiber/v2"
)

type heathCheckHandler struct {
	App    string `json:"app"`
	Status string `json:"status"`
}

func (s *server) HealthCheck(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(&heathCheckHandler{
		App:    s.cfg.App.Name,
		Status: "OK",
	})
}
