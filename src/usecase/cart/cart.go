package cart

import (
	"context"
	"github.com/itsmeadi/cart/src/domain/repositories"
	"github.com/itsmeadi/cart/src/entities/constants"
	"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/usecase/definitions"
)

type CartUseCase struct {
	CartRepo          repositories.Cart
	CartItemsRepo     repositories.CartItems
	CartXCartItemRepo repositories.CartXCartItems
	Products          repositories.Product
}

var _ definitions.Cart = &CartUseCase{}

func InitCartUseCase(uc CartUseCase) CartUseCase {
	return uc
}

//func (cart *CartUseCase) AddToCart(ctx context.Context, userId, productId, qty int64) error {
//
//	cartItems, err := cart.CartXCartItemRepo.GetCartItemsByUserAndProductId(ctx, userId, productId, constants.StatusActive, constants.StatusActive)
//	if err != nil {
//		return err
//	}
//	var cartId int64
//	if len(cartItems) == 0 {
//		cartId, err = cart.CartRepo.CreateCart(ctx, models.Cart{
//			UserID: userId,
//			Status: constants.StatusActive,
//		})
//		if err != nil {
//			return err
//		}
//	} else {
//		cartId = cartItems[0].CartID
//	}
//
//	//check if product already exist in cart, id it does, then update the quantity
//	for _, item := range cartItems {
//		if item.ProductID == productId {
//			_, err := cart.CartItemsRepo.UpdateCartQty(ctx, qty+item.Qty, item.CartID, item.ProductID)
//			if err != nil {
//				return err
//			}
//			return nil
//		}
//	}
//
//	_, err = cart.CartItemsRepo.InsertItemInCart(ctx, models.CartItems{
//		ProductID: productId,
//		Qty:       qty,
//		Status:    constants.StatusActive,
//		CartID:    cartId,
//	})
//	return err
//}

//func (cart *CartUseCase) AddToCart(ctx context.Context, userId, productId, qty int64) error {
//
//	var item models.CartXCartItems
//	cartItems, err := cart.CartXCartItemRepo.GetCartItemsByUserAndProductId(ctx, userId, productId, constants.StatusActive, constants.StatusActive)
//	if err != nil {
//		return err
//	}
//	var cartId int64
//	if len(cartItems) == 0 {
//		cartId, err = cart.CartRepo.CreateCart(ctx, models.Cart{
//			UserID: userId,
//			Status: constants.StatusActive,
//		})
//		if err != nil {
//			return err
//		}
//	} else {
//		cartId = cartItems[0].CartID
//		item = cartItems[0]
//	}
//
//	//check if product already exist in cart, id it does, then update the quantity
//	if item.ProductID != 0 {
//		if item.ProductID == productId {
//			_, err = cart.CartItemsRepo.UpdateCartQty(ctx, qty+item.Qty, item.CartID, item.ProductID)
//		}
//	} else {
//		_, err = cart.CartItemsRepo.InsertItemInCart(ctx, models.CartItems{
//			ProductID: productId,
//			Qty:       qty,
//			Status:    constants.StatusActive,
//			CartID:    cartId,
//		})
//	}
//	return err
//}

func (cart *CartUseCase) AddToCart(ctx context.Context, userId, productId, qty int64) error {

	product, err := cart.Products.GetProductDetailById(ctx, productId)

	if err != nil || product.ID == 0 {
		return constants.ERRPRODUCTUNAVAILABLE //TODO
	}

	crt, err := cart.CartRepo.GetCart(ctx, userId)
	if err != nil {
		return err
	}
	var cartId int64
	if len(crt) == 0 {
		cartId, err = cart.CartRepo.CreateCart(ctx, models.Cart{
			UserID: userId,
			Status: constants.StatusActive,
		})
		if err != nil {
			return err
		}
	} else {
		cartId = crt[0].ID
	}

	cartItem, err := cart.CartItemsRepo.GetProductFromCart(ctx, cartId, productId, constants.StatusActive)
	if err != nil {
		return err
	}
	//check if product already exist in cart, id it does, then update the quantity
	if len(cartItem) != 0 {
		item := cartItem[0]
		if item.ProductID == productId {
			_, err = cart.CartItemsRepo.UpdateCartQty(ctx, qty+item.Qty, item.CartID, item.ProductID)
		}
	} else {
		_, err = cart.CartItemsRepo.InsertItemInCart(ctx, models.CartItems{
			ProductID: productId,
			Qty:       qty,
			Status:    constants.StatusActive,
			CartID:    cartId,
		})
	}
	return err
}
func (cart *CartUseCase) UpdateCart(ctx context.Context, userId, productId, qty int64) error {

	var item models.CartXCartItems
	cartItems, err := cart.CartXCartItemRepo.GetCartItemsByUserAndProductId(ctx, userId, productId, constants.StatusActive, constants.StatusActive)
	if err != nil {
		return err
	}
	var cartId int64
	if len(cartItems) == 0 {
		cartId, err = cart.CartRepo.CreateCart(ctx, models.Cart{
			UserID: userId,
			Status: constants.StatusActive,
		})
		if err != nil {
			return err
		}
	} else {
		cartId = cartItems[0].CartID
		item = cartItems[0]
	}

	//check if product already exist in cart, id it does, then update the quantity
	if item.ProductID != 0 {
		if item.ProductID == productId {
			_, err = cart.CartItemsRepo.UpdateCartQty(ctx, qty, item.CartID, item.ProductID)
		}
	} else {
		_, err = cart.CartItemsRepo.InsertItemInCart(ctx, models.CartItems{
			ProductID: productId,
			Qty:       qty,
			Status:    constants.StatusActive,
			CartID:    cartId,
		})
	}
	return err
}

func (cart *CartUseCase) RemoveProductFromCart(ctx context.Context, productId, userId int64) error {
	crt, err := cart.CartRepo.GetCart(ctx, userId)
	if err != nil || len(crt) == 0 {
		return err
	}
	productId = 4
	_, err = cart.CartItemsRepo.UpdateCartQty(ctx, 0, crt[0].ID, productId)
	return err
}

func (cart *CartUseCase) GetCart(ctx context.Context, userId int64) (models.CartDetail, error) {

	var cDetail models.CartDetail
	cartItems, err := cart.CartXCartItemRepo.GetCartItemsByUser(ctx, userId, constants.StatusActive, constants.StatusActive)
	if err != nil {
		return cDetail, err
	}
	if len(cartItems) == 0 {
		return cDetail, nil
	}
	cDetail.CartID = cartItems[0].CartID
	products := make([]models.CartProduct, 0)
	for _, item := range cartItems {
		if item.Qty <= 0 {
			continue
		}
		var cartProd models.CartProduct
		prod, err := cart.Products.GetProductDetailById(ctx, item.ProductID)
		if err != nil {
			continue
		}
		cartProd.Product = prod
		cartProd.Qty = item.Qty
		cartProd.TotalPrice = float64(item.Qty) * prod.Price
		products = append(products, cartProd)

		cDetail.TotalPrice += cartProd.TotalPrice
	}

	cDetail.UserID = userId
	cDetail.Products = products
	return cDetail, err
}
