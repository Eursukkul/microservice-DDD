package request

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	contextWrapperService interface {
		Bind(data any) error
	}

	contexWrapper struct {
		Context   *fiber.Ctx
		validator *validator.Validate
	}
)

func ContextWrapper(ctx *fiber.Ctx) contextWrapperService {
	return &contexWrapper{
		Context:   ctx,
		validator: validator.New(),
	}
}

func (c *contexWrapper) Bind(data any) error {
	if err := c.Context.BodyParser(data); err != nil {
		log.Printf("Error: Bind data failed: %s", err.Error())
	}

	if err := c.validator.Struct(data); err != nil {
		log.Printf("Error: Bind data failed: %s", err.Error())
	}

	return nil
}
