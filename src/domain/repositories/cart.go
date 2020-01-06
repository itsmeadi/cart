package repositories

import (
	"context"
	"github.com/itsme/cart/src/entities/models"
)

type Cart interface {
	CreateCart(ctx context.Context, cart models.Cart) (int64, error)
	GetCart(ctx context.Context, userId int64) ([]models.Cart, error)

}
