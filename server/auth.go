package server

import (
	"gitlab.com/chalermphanEur/modules/auth/authHandler"

	"gitlab.com/chalermphanEur/modules/auth/authRepository"
	"gitlab.com/chalermphanEur/modules/auth/authUsecase"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	handler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	_ = handler

	auth := s.app.Group("/auth_v1")

	//Health check
	auth.Get("/healthcheck", s.HealthCheck)
}
