package elasticSearchHandler

import (
	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/elasticSearch/elasticSearchUsecase"
)

type (
	IElasticSearchHandler interface {
	}

	elasticSearchHandler struct {
		cfg                  *config.Config
		elasticSearchUsecase elasticSearchUsecase.IElasticSearchUsecase
	}
)

func NewElasticSearchHandler(cfg *config.Config, elasticSearchUsecase elasticSearchUsecase.IElasticSearchUsecase) IElasticSearchHandler {
	return &elasticSearchHandler{
		cfg:                  cfg,
		elasticSearchUsecase: elasticSearchUsecase,
	}
}
