package mysql

import (
	"context"
	"database/sql"
	"github.com/itsme/cart/src/domain/repositories"

	"github.com/itsme/cart/src/entities/models"
)

var _ repositories.CartItems = &DB{}
var _ repositories.Cart = &DB{}

func (Conn *DB) CreateCart(ctx context.Context, cart models.Cart) (int64, error) {
	res, err := Conn.queries.CreateCart.Exec(cart.UserID, cart.Status)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetCart(ctx context.Context, userId int64) ([]models.Cart, error) {
	var cart []models.Cart
	err := Conn.queries.GetCart.SelectContext(ctx, &cart, userId)
	return cart, err
}

func (Conn *DB) InsertItemInCart(ctx context.Context, item models.CartItems) (int64, error) {
	res, err := Conn.queries.InsertItem.Exec(item.CartID, item.ProductID, item.Qty, item.Status)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}

func (Conn *DB) GetItemsInCart(ctx context.Context, cartId, status, limit int64) ([]models.CartItems, error) {
	var items []models.CartItems
	err := Conn.queries.GetItemsInCart.SelectContext(ctx, &items, cartId, status, limit)
	return items, err
}

func (Conn *DB) GetProductFromCart(ctx context.Context, cartId, productId int64, status int) ([]models.CartItems, error) {
	var items []models.CartItems
	err := Conn.queries.GetProductFromCart.SelectContext(ctx, &items, cartId, productId, status)
	return items, err
}

func (Conn *DB) UpdateCartQty(ctx context.Context, qty, cartId, productId int64) (int64, error) {
	res, err := Conn.queries.UpdateCartQty.Exec(qty, cartId, productId)
	if err != nil {
		return 0, err
	}
	rowAff, _ := res.RowsAffected()
	return rowAff, err
}

func (Conn *DB) UpdateCartItemStatus(ctx context.Context, status, cartId, productId int64) (int64, error) {
	res, err := Conn.queries.UpdateCartItemStatus.Exec(status, cartId, productId)
	if err != nil {
		return 0, err
	}
	rowAff, _ := res.RowsAffected()
	return rowAff, err
}

func (Conn *DB) GetCartItemsByUser(ctx context.Context, userId int64, cartStatus, itemStatus int) ([]models.CartXCartItems, error) {
	var items []models.CartXCartItems
	err := Conn.queries.GetCartItemsByUser.SelectContext(ctx, &items, userId, cartStatus, itemStatus)
	return items, err
}

func (Conn *DB) GetCartItemsByUserAndProductId(ctx context.Context, userId, productId int64, cartStatus, itemStatus int) ([]models.CartXCartItems, error) {
	var items []models.CartXCartItems
	err := Conn.queries.GetCartItemsByUserAndProductId.SelectContext(ctx, &items, userId, productId, cartStatus, itemStatus)
	return items, err
}

func (Conn *DB) GetUserBySub(ctx context.Context, sub string) (models.User, error) {
	var items models.User
	err := Conn.queries.GetUserBySub.GetContext(ctx, &items, sub)
	if err == sql.ErrNoRows {
		return items, nil
	}
	return items, err
}

func (Conn *DB) AddUser(ctx context.Context, user models.User) (int64, error) {
	res, err := Conn.queries.AddUser.Exec(user.Sub)
	if err != nil {
		return 0, err
	}
	lID, _ := res.LastInsertId()
	return lID, err
}
