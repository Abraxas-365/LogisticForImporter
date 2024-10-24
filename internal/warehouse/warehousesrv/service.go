package warehousesrv

import (
	"context"

	"github.com/Abraxas-365/cabo/internal/warehouse"
	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Service struct {
	repo warehouse.Repository
}

func New(repo warehouse.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SaveWarehouse(ctx context.Context, u *warehouse.Warehouse) (*warehouse.Warehouse, error) {
	return s.repo.SaveWarehouse(ctx, u)
}

func (s *Service) GetWarehouseById(ctx context.Context, id int) (*warehouse.Warehouse, error) {
	return s.repo.GetWarehouseById(ctx, id)
}

func (s *Service) UpdateWarehouse(ctx context.Context, u *warehouse.Warehouse) (*warehouse.Warehouse, error) {
	return s.repo.UpdateWarehouse(ctx, u)
}

func (s *Service) GetAllWarehouses(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[warehouse.Warehouse], error) {
	return s.repo.GetAllWarehouses(ctx, page, pageSize)
}
