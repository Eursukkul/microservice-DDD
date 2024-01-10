package usersHandler

import (
	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/users/usersUsecase"
)

type (
	IUsersQueueHandlerService interface {
	}

	usersQueueHandler struct {
		cfg          *config.Config
		usersUsecase usersUsecase.IUsersUsecaseService
	}
)

func NewUserQueuesHandlerService(cfg *config.Config, usersUsecase usersUsecase.IUsersUsecaseService) IUsersQueueHandlerService {
	return &usersQueueHandler{
		cfg:          cfg,
		usersUsecase: usersUsecase,
	}
}
