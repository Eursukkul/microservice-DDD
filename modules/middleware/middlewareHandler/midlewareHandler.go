package middlewareHandler

import (
	
	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/middleware/middlewareUsecase"
	
)

type (
	IMiddleHandlerService interface {
	}

	middlewareHandlerService struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.IMiddlewareUsercase
	}
)

func NewMiddlewareHandlerService(cfg *config.Config, middlewareUsecase middlewareUsecase.IMiddlewareUsercase) IMiddleHandlerService {
	return &middlewareHandlerService{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}