package product_usecase

import (
	"context"
	"github.com/itsme/cart/src/domain/repositories"
	"github.com/itsme/cart/src/entities/models"
)

type Product struct {
	ProductRepo repositories.Product
}

func InitProductUseCase(repo repositories.Product) Product {
	return Product{
		ProductRepo: repo,
	}
}

func (product *Product) GetProductDetailById(ctx context.Context, id int64) (models.Product, error) {
	return product.ProductRepo.GetProductDetailById(ctx, id)
}
