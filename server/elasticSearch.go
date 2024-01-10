package server

import (
	"gitlab.com/chalermphanEur/modules/elasticSearch/elasticSearchHandler"
	"gitlab.com/chalermphanEur/modules/elasticSearch/elasticSearchRepository"
	"gitlab.com/chalermphanEur/modules/elasticSearch/elasticSearchUsecase"
)

func (s *server) elasticSearchService() {
	repo := elasticSearchRepository.NewElasticSearchRepository(s.elastic)
	usecase := elasticSearchUsecase.NewElasticSearchUsecase(repo)
	handler := elasticSearchHandler.NewElasticSearchHandler(s.cfg, usecase)

	_ = handler

	elasticSearch := s.app.Group("/elasticSearch_v1")

	//Health check
	elasticSearch.Get("/healthcheck", s.HealthCheck)
}
