package warehouse

import (
	"context"

	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Repository interface {
	SaveWarehouse(ctx context.Context, u *Warehouse) (*Warehouse, error)
	GetWarehouseById(ctx context.Context, id int) (*Warehouse, error)
	UpdateWarehouse(ctx context.Context, u *Warehouse) (*Warehouse, error)
	GetAllWarehouses(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[Warehouse], error)
}
