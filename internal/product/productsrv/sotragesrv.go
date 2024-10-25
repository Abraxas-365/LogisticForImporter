package productsrv

import (
	"context"

	"github.com/Abraxas-365/cabo/internal/product"
	"github.com/Abraxas-365/toolkit/pkg/database"
	"github.com/Abraxas-365/toolkit/pkg/errors"
)

func (s *Service) GetUserProductsInStorage(ctx context.Context, userID, warehouseID, page, pageSize int) (database.PaginatedRecord[product.Product], error) {
	return s.repo.GetAllUserProductsByStatusAndWarehouse(ctx, product.InWarehouse, userID, warehouseID, page, pageSize)
}

func (s *Service) ProductArrivedToStorage(ctx context.Context, productID int) error {
	p, err := s.GetProductById(ctx, productID)
	if err != nil {
		return err
	}

	if p.Status != product.TransitToWarehouse {
		return errors.ErrUnexpected("Product is not in transit")
	}

	p.Status = product.InWarehouse

	if _, err := s.UpdateProduct(ctx, p); err != nil {
		return err
	}

	return nil
}
