package elasticSearchRepository

import "github.com/elastic/go-elasticsearch/v8"

type (
	IElasticSearchRepository interface {
	}

	elasticSearchRepository struct {
		elasticsearch *elasticsearch.Client
	}
)

func NewElasticSearchRepository(elasticsearch *elasticsearch.Client) IElasticSearchRepository {
	return &elasticSearchRepository{elasticsearch}
}
