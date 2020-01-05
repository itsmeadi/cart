package definitions

import (
	"context"
	"github.com/itsme/nuc/src/entities/models"
)

type Cart interface {
	UpdateCart(ctx context.Context, userId, productId, qty int64) error
	AddToCart(ctx context.Context, userId, productId, qty int64) error
	GetCart(ctx context.Context, userId int64) (models.CartDetail, error)
	RemoveProductFromCart(ctx context.Context, productId, userId int64) error
}
