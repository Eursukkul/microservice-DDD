package authUsecase

import "gitlab.com/chalermphanEur/modules/auth/authRepository"

type (
	IAuthUsecaseService interface {
	}

	authUsecase struct {
		authRepository authRepository.IAuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authRepository.IAuthRepositoryService) IAuthUsecaseService {
	return &authUsecase{authRepository}
}
