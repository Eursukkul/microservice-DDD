package authHandler

import (
	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/auth/authUsecase"
)

type (
	IAuthHttpHandlerService interface {
	}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.IAuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authUsecase.IAuthUsecaseService) IAuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}
