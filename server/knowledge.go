package server

import (
	"gitlab.com/chalermphanEur/modules/knowledge/knowledgeHandler"
	"gitlab.com/chalermphanEur/modules/knowledge/knowledgeRepository"
	"gitlab.com/chalermphanEur/modules/knowledge/knowledgeUsecase"
)

func (s *server) kbmService() {
	repo := knowledgeRepository.NewKnowledgeRepository(s.mongoDb)
	usecase := knowledgeUsecase.NewKnowleadgeusecase(repo)
	handler := knowledgeHandler.NewKnowledgeHandlerService(s.cfg, usecase)

	_ = handler

	kbm := s.app.Group("/kbm_v1")

	//Health check
	kbm.Get("/healthcheck", s.HealthCheck)
}
