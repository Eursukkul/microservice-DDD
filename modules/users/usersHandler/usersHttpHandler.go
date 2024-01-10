package usersHandler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/chalermphanEur/config"

	"gitlab.com/chalermphanEur/modules/users"
	"gitlab.com/chalermphanEur/modules/users/usersUsecase"
)

type (
	IUsersHandlerHttpService interface {
		UserSearch(c *fiber.Ctx) error
		UserRegister(c *fiber.Ctx) error
	}

	usersHttpHandler struct {
		cfg          *config.Config
		usersUsecase usersUsecase.IUsersUsecaseService
	}
)

func NewUsersHandlerService(cfg *config.Config, usersUsecase usersUsecase.IUsersUsecaseService) IUsersHandlerHttpService {
	return &usersHttpHandler{
		cfg:          cfg,
		usersUsecase: usersUsecase,
	}
}

func (h *usersHttpHandler) UserSearch(c *fiber.Ctx) error {

	result, err := h.usersUsecase.UserSearch(c.Params("employeeId"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    result,
	})

}

func (h *usersHttpHandler) UserRegister(c *fiber.Ctx) error {

	user := new(users.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if h.usersUsecase.FindUser(c.Context(), user.EmployeeId) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Employee ID already exists",
		})
	}

	err := h.usersUsecase.UserRegister(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
	})
}

