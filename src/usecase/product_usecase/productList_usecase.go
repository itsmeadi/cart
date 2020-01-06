package product_usecase

import (
	"context"
	"github.com/itsme/cart/src/domain/repositories"
	"github.com/itsme/cart/src/entities/models"
)

type ProductList struct {
	ProductListRepo repositories.ProductResponseList
}


func InitProductListUseCase(repo repositories.ProductResponseList) ProductList {
	return ProductList{
		ProductListRepo: repo,
	}
}

func (product *ProductList) GetProductListByCategoryId(ctx context.Context, id int64) (models.ProductListResponse, error) {
	return product.ProductListRepo.GetProductListByCategoryId(ctx, id)
}


func (product *ProductList) GetProductArrByCategoryId(ctx context.Context, id int64) ([]models.Product, error) {
	return product.ProductListRepo.GetProductArrByCategoryId(ctx, id)
}


