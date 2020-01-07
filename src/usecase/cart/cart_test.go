package cart

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/itsmeadi/cart/src/entities/constants"
	"github.com/itsmeadi/cart/src/entities/models"
	"github.com/itsmeadi/cart/src/interfaces/mock_repositories"
	"log"
	"testing"
)

func TestAddToCart(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic in f", r)
		}
	}()
	type Args struct {
		ctx                            context.Context
		userId, productId, qty, cartId int64
	}

	tests := []struct {
		name                   string
		args                   Args
		MockGetCart            func(mock *mock_repositories.MockCart, args Args)
		MockGetProductFromCart func(mock *mock_repositories.MockCartItems, args Args)
		MockInsertItemInCart   func(mock *mock_repositories.MockCartItems, args Args)

		wantErr bool
	}{

		{
			"success",
			Args{
				ctx:       context.Background(),
				userId:    1,
				productId: 2,
				qty:       3,
				cartId:    5,
			},
			func(mock *mock_repositories.MockCart, args Args) {
				mock.EXPECT().GetCart(gomock.Any(), args.userId).
					Return(
						[]models.Cart{
							{
								ID:     args.cartId,
								UserID: args.userId,
								Status: 1,
							},
						}, nil)
			},

			func(mock *mock_repositories.MockCartItems, args Args) {
				mock.EXPECT().GetProductFromCart(gomock.Any(), args.cartId, args.productId, constants.StatusActive).
					Return(
						[]models.CartItems{}, nil)

			},
			func(mock *mock_repositories.MockCartItems, args Args) {
				mock.EXPECT().InsertItemInCart(gomock.Any(), models.CartItems{
					CartID:    5,
					ProductID: args.productId,
					Qty:       args.qty,
					Status:    constants.StatusActive,
				}).Return(int64(1), nil)
			},
			false,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range tests {

		mockCart := mock_repositories.NewMockCart(ctrl)
		mockCartItemsRepo := mock_repositories.NewMockCartItems(ctrl)
		mockCartXCartItemRepo := mock_repositories.NewMockCartXCartItems(ctrl)
		mockProducts := mock_repositories.NewMockProduct(ctrl)

		cartUseCaseMock := InitCartUseCase(CartUseCase{
			CartRepo:          mockCart,
			CartItemsRepo:     mockCartItemsRepo,
			CartXCartItemRepo: mockCartXCartItemRepo,
			Products:          mockProducts,
		})
		tt.MockGetCart(mockCart, tt.args)
		tt.MockInsertItemInCart(mockCartItemsRepo, tt.args)
		tt.MockGetProductFromCart(mockCartItemsRepo, tt.args)
		err := cartUseCaseMock.AddToCart(tt.args.ctx, tt.args.userId, tt.args.productId, tt.args.qty)
		if (err == nil) == tt.wantErr {
			t.Errorf("Case Failed=%+v", tt.name)
		}

	}
}
