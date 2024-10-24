package product

import (
	"context"

	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Repository interface {
	SaveProduct(ctx context.Context, u *Product) (*Product, error)
	GetProductById(ctx context.Context, id int) (*Product, error)
	UpdateProduct(ctx context.Context, u *Product) (*Product, error)
	GetAllUserProducts(ctx context.Context, userID int, page int, pageSize int) (database.PaginatedRecord[Product], error)
}
