package authRepository

import (
	"gorm.io/gorm"
)

type (
	IAuthRepositoryService interface {
	}

	authRepository struct {
		db *gorm.DB
	}
)

func NewAuthRepository(db *gorm.DB) IAuthRepositoryService {
	return &authRepository{db}
}
