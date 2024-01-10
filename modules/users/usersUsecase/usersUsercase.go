package usersUsecase

import (
	"context"
	"time"

	"gitlab.com/chalermphanEur/modules/users"
	"gitlab.com/chalermphanEur/modules/users/usersReository"
)

type (
	IUsersUsecaseService interface {
		UserSearch(employeeId string) (*users.UserProfile, error)
		UserRegister(ctx context.Context, user *users.User) error
		FindUser(ctx context.Context, employeeId string) bool
	}

	usersUsecase struct {
		usersReository usersReository.IUsersRepositoryService
	}
)

func NewUsersUsecase(usersReository usersReository.IUsersRepositoryService) IUsersUsecaseService {
	return &usersUsecase{
		usersReository: usersReository,
	}
}

func (u *usersUsecase) UserSearch(employeeId string) (*users.UserProfile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	userProfile, err := u.usersReository.GetUserProfile(ctx, employeeId)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}

func (u *usersUsecase) UserRegister(ctx context.Context, user *users.User) error {

	err := u.usersReository.InsertUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *usersUsecase) FindUser(ctx context.Context, employeeId string) bool {

	result := u.usersReository.FindUser(ctx, employeeId)
	if result {
		return true
	}

	return false

}

func (u *usersUsecase) AuthUrl(username, password string) error {

	return nil
}
