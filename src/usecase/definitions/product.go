package definitions

import (
	"context"
	"github.com/itsmeadi/cart/src/entities/models"
)

type Product interface {
	GetProductDetailById(ctx context.Context, id int64) (models.Product, error)
}

type ProductByCategory interface {
	GetProductListByCategoryId(ctx context.Context, id int64) (models.ProductListResponse, error)
	GetProductArrByCategoryId(ctx context.Context, id int64) ([]models.Product, error)
}
