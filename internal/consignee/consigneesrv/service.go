package consigneesrv

import (
	"context"

	"github.com/Abraxas-365/cabo/internal/consignee"
	"github.com/Abraxas-365/cabo/pkg/apis"
	"github.com/Abraxas-365/toolkit/pkg/database"
	"github.com/Abraxas-365/toolkit/pkg/errors"
)

type Service struct {
	repo    consignee.Repository
	userSrv apis.UserServicer
}

func NewService(repo consignee.Repository, userSrv apis.UserServicer) *Service {
	return &Service{
		repo:    repo,
		userSrv: userSrv,
	}
}

// CreateConsignatario creates a new consignatario
func (s *Service) CreateConsignatario(ctx context.Context, c *consignee.Consignee) (*consignee.Consignee, error) {

	if _, err := s.userSrv.GetUser(ctx, c.UserID); err != nil {
		return nil, err
	}

	// Validate consignatario if needed
	if err := c.DocumentType.IsValid(); err != nil {
		return nil, errors.ErrBadRequest(err.Error())
	}

	nc, err := s.repo.SaveConsignatario(ctx, c)
	if err != nil {
		return nil, errors.ErrDatabase(err.Error())
	}

	return nc, nil
}

// GetConsignatario retrieves a consignatario by ID
func (s *Service) GetConsignatario(ctx context.Context, id int) (*consignee.Consignee, error) {
	c, err := s.repo.GetConsignatarioById(ctx, id)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, errors.ErrNotFound("Consignatario not found")
		}
		return nil, errors.ErrDatabase(err.Error())
	}

	return c, nil
}

// UpdateConsignatario updates an existing consignatario
func (s *Service) UpdateConsignatario(ctx context.Context, c *consignee.Consignee) (*consignee.Consignee, error) {
	if _, err := s.userSrv.GetUser(ctx, c.UserID); err != nil {
		return nil, err
	}

	// Validate consignatario if needed
	if err := c.DocumentType.IsValid(); err != nil {
		return nil, err
	}

	updatedConsignatario, err := s.repo.UpdateConsignatario(ctx, c)
	if err != nil {
		return nil, err
	}

	return updatedConsignatario, nil
}

// GetAllByUserId retrieves all consignatarios for a given user ID with pagination
func (s *Service) GetAllByUserId(ctx context.Context, userId int, pageNumber, pageSize int) (database.PaginatedRecord[consignee.Consignee], error) {
	paginatedConsignatarios, err := s.repo.GetAllByUserId(ctx, userId, pageNumber, pageSize)
	if err != nil {
		return paginatedConsignatarios, err
	}

	return paginatedConsignatarios, nil
}
