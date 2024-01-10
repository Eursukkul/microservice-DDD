package knowledgeRepository

import "go.mongodb.org/mongo-driver/mongo"

type (
	IKnowledgeRepositoryService interface {
	}

	knowledgeRepository struct {
		db *mongo.Client
	}
)

func NewKnowledgeRepository(db *mongo.Client) IKnowledgeRepositoryService {
	return &knowledgeRepository{db}
}
