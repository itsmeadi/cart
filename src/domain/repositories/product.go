package repositories

import (
	"context"
	"github.com/itsme/nuc/src/entities/models"
)

type Product interface {
	GetProductDetailById(ctx context.Context, id int64)(models.Product, error)
}
