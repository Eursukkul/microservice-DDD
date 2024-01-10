package usersReository

import (
	"context"
	"errors"

	"time"

	"gitlab.com/chalermphanEur/modules/users"
	"gorm.io/gorm"
)

type (
	IUsersRepositoryService interface {
		FindUser(ctx context.Context, employeeId string) bool
		InsertUser(ctx context.Context, user *users.User) error
		GetUserProfile(ctx context.Context, employeeId string) (*users.UserProfile, error)
	}

	usersRepository struct {
		db *gorm.DB
	}
)

func NewUsersRepository(db *gorm.DB) IUsersRepositoryService {
	return &usersRepository{db}
}

func (u *usersRepository) FindUser(ctx context.Context, employeeId string) bool {

	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	db := u.db.Where("employee_id = ?", employeeId).First(&users.User{})
	if err := db.Error; err != nil {
		return false
	}
	return true
}

func (u *usersRepository) InsertUser(ctx context.Context, user *users.User) error {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	db := u.db.Create(&user)
	if err := db.Error; err != nil {
		return errors.New("insert user failed")
	}

	return nil
}

func (u *usersRepository) GetUserProfile(ctx context.Context, employeeId string) (*users.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	userProfile := new(users.UserProfile)
	db := u.db.Where("employee_id = ?", employeeId).First(&userProfile)
	if err := db.Error; err != nil {
		return nil, errors.New("user not found")
	}

	return userProfile, nil
}
