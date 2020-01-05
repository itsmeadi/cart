package repositories

import (
	"context"
	"github.com/itsme/nuc/src/entities/models"
)

type CartXCartItems interface {
	GetCartItemsByUserAndProductId(ctx context.Context, userId, productId int64, cartStatus, itemStatus int) ([]models.CartXCartItems, error)
	GetCartItemsByUser(ctx context.Context, userId int64, cartStatus, itemStatus int) ([]models.CartXCartItems, error)
}
