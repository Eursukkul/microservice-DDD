package elasticSearchUsecase

import "gitlab.com/chalermphanEur/modules/elasticSearch/elasticSearchRepository"

type (
	IElasticSearchUsecase interface {
	}

	elasticSearchUsecase struct {
		elasticSearchRepository elasticSearchRepository.IElasticSearchRepository
	}
)

func NewElasticSearchUsecase(elasticSearchRepository elasticSearchRepository.IElasticSearchRepository) IElasticSearchUsecase {
	return &elasticSearchUsecase{
		elasticSearchRepository: elasticSearchRepository,
	}
}
