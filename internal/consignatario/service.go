package consignatario

import (
	"context"

	"github.com/Abraxas-365/toolkit/pkg/database"
	"github.com/Abraxas-365/toolkit/pkg/errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateConsignatario creates a new consignatario
func (s *Service) CreateConsignatario(ctx context.Context, consignatario *Consignatario) (*Consignatario, error) {
	// Validate consignatario if needed
	if err := consignatario.DocumentType.IsValid(); err != nil {
		return nil, errors.ErrBadRequest(err.Error())
	}

	nc, err := s.repo.SaveConsignatario(ctx, consignatario)
	if err != nil {
		return nil, errors.ErrDatabase(err.Error())
	}

	return nc, nil
}

// GetConsignatario retrieves a consignatario by ID
func (s *Service) GetConsignatario(ctx context.Context, id int) (*Consignatario, error) {
	consignatario, err := s.repo.GetConsignatarioById(ctx, id)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, errors.ErrNotFound("Consignatario not found")
		}
		return nil, errors.ErrDatabase(err.Error())
	}

	return consignatario, nil
}

// UpdateConsignatario updates an existing consignatario
func (s *Service) UpdateConsignatario(ctx context.Context, consignatario *Consignatario) (*Consignatario, error) {
	// Validate consignatario if needed
	if err := consignatario.DocumentType.IsValid(); err != nil {
		return nil, err
	}

	updatedConsignatario, err := s.repo.UpdateConsignatario(ctx, consignatario)
	if err != nil {
		return nil, err
	}

	return updatedConsignatario, nil
}

// GetAllByUserId retrieves all consignatarios for a given user ID with pagination
func (s *Service) GetAllByUserId(ctx context.Context, userId int, pageNumber, pageSize int) (database.PaginatedRecord[Consignatario], error) {
	paginatedConsignatarios, err := s.repo.GetAllByUserId(ctx, userId, pageNumber, pageSize)
	if err != nil {
		return paginatedConsignatarios, err
	}

	return paginatedConsignatarios, nil
}
