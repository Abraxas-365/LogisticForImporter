package productsrv

import (
	"context"
	"fmt"

	"github.com/Abraxas-365/cabo/internal/product"
	"github.com/Abraxas-365/toolkit/pkg/database"
	"github.com/Abraxas-365/toolkit/pkg/errors"
	"github.com/Abraxas-365/toolkit/pkg/s3client"
	"github.com/google/uuid"
)

type Service struct {
	repo product.Repository
	s3   s3client.Client
}

func (s *Service) SaveProduct(ctx context.Context, u *product.Product) (*product.Product, error) {
	return s.repo.SaveProduct(ctx, u)
}
func (s *Service) GetProductById(ctx context.Context, id int) (*product.Product, error) {
	return s.repo.GetProductById(ctx, id)
}

func (s *Service) UpdateProduct(ctx context.Context, u *product.Product) (*product.Product, error) {
	return s.repo.UpdateProduct(ctx, u)
}

func (s *Service) GeneratePresignedPutURL(ctx context.Context, userID int) (string, error) {
	key := fmt.Sprintf("user/%d/invoice/%s", userID, uuid.New().String())
	url, err := s.s3.GeneratePresignedPutURL(key, 60)
	if err != nil {
		return "", errors.ErrUnexpected(err.Error())
	}

	return url, nil
}

func (s *Service) GetAllUserProducts(ctx context.Context, userID, page, pageSize int) (database.PaginatedRecord[product.Product], error) {
	return s.repo.GetAllUserProducts(ctx, userID, page, pageSize)
}
