package server

import (
	"gitlab.com/chalermphanEur/modules/users/usersHandler"
	"gitlab.com/chalermphanEur/modules/users/usersReository"
	"gitlab.com/chalermphanEur/modules/users/usersUsecase"
)

func (s *server) userService() {
	repo := usersReository.NewUsersRepository(s.db)
	usecase := usersUsecase.NewUsersUsecase(repo)
	handler := usersHandler.NewUsersHandlerService(s.cfg, usecase)
	queueHandler := usersHandler.NewUserQueuesHandlerService(s.cfg, usecase)

	users := s.app.Group("/users_v1")

	//Health check
	users.Get("/healthcheck", s.HealthCheck)
	users.Get("/:employeeId", handler.UserSearch)
	users.Post("/register", handler.UserRegister)

	//Queue
	_ = queueHandler
}
