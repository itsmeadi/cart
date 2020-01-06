package repositories

import (
	"context"
	"github.com/itsme/cart/src/entities/models"
)

type CartItems interface {
	InsertItemInCart(ctx context.Context, item models.CartItems) (int64, error)
	GetItemsInCart(ctx context.Context, cartId, status, limit int64) ([]models.CartItems, error)
	GetProductFromCart(ctx context.Context, cartId, productId int64, status int) ([]models.CartItems, error)
	UpdateCartQty(ctx context.Context, qty, cartId, productId int64) (int64, error)
	UpdateCartItemStatus(ctx context.Context, status, cartId, productId int64) (int64, error)
}
