package repositories

import (
	"context"
	"github.com/itsme/nuc/src/entities/models"
)

type ProductResponseList interface {
	GetProductListByCategoryId(ctx context.Context, categoryId int64)(models.ProductListResponse, error)
	GetProductArrByCategoryId(ctx context.Context, id int64) ([]models.Product, error)
}
