package knowledgeUsecase

import "gitlab.com/chalermphanEur/modules/knowledge/knowledgeRepository"

type (
	IKnowledgeUsecasecaseService interface {
	}

	knowledgeUsecase struct {
		knowledgeRepository knowledgeRepository.IKnowledgeRepositoryService
	}
)

func NewKnowleadgeusecase(knowledgeRepository knowledgeRepository.IKnowledgeRepositoryService) IKnowledgeUsecasecaseService {
	return &knowledgeUsecase{
		knowledgeRepository: knowledgeRepository,
	}
}
