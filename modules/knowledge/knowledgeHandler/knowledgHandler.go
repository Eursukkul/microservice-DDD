package knowledgeHandler

import (
	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/knowledge/knowledgeUsecase"
)

type (
	IKnowledgeHandlerService interface {
	}

	knowledgeHandlerService struct {
		cfg               *config.Config
		knowledgeUsercase knowledgeUsecase.IKnowledgeUsecasecaseService
	}
)

func NewKnowledgeHandlerService(cfg *config.Config, knowledgeUsecase knowledgeUsecase.IKnowledgeUsecasecaseService) IKnowledgeHandlerService {
	return &knowledgeHandlerService{
		cfg:               cfg,
		knowledgeUsercase: knowledgeUsecase,
	}
}
