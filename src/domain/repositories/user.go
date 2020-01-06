package repositories

import (
	"context"
	"github.com/itsme/cart/src/entities/models"
)

type User interface {
	GetUserBySub(ctx context.Context, sub string) (models.User, error)
	AddUser(ctx context.Context, user models.User) (int64, error)
}
