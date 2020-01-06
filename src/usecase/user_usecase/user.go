package user_usecase

import (
	"context"
	"github.com/itsme/cart/src/domain/repositories"
	"github.com/itsme/cart/src/entities/models"
)

type User struct {
	UserRepo repositories.User
}

func InitUsecase(repo repositories.User)*User{
	return &User{
		UserRepo:repo,
	}
}
func (userUC *User) GetUserBySub(ctx context.Context, sub string) (models.User, error) {
	return userUC.UserRepo.GetUserBySub(ctx, sub)
}

func (userUC *User) AddUser(ctx context.Context, user models.User) (int64, error) {
	return userUC.UserRepo.AddUser(ctx, user)
}
