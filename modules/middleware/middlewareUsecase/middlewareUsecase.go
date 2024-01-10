package middlewareUsecase

import (
	
	"gitlab.com/chalermphanEur/modules/middleware/middlewareRepository"
	
)

type (
	IMiddlewareUsercase interface {
	}

	middlewareUsercase struct {
		middlewareRepository middlewareRepository.IMiddlewareRepository
	}
)

func NewMiddlewareUsercase(middlewareRepository middlewareRepository.IMiddlewareRepository) IMiddlewareUsercase {
	return &middlewareUsercase{
		middlewareRepository: middlewareRepository,
	}
}

