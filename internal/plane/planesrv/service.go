package planesrv

import (
	"context"
	"time"

	"github.com/Abraxas-365/cabo/internal/plane"
	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Service struct {
	db plane.Repository
}

func New(db plane.Repository) *Service {
	return &Service{db: db}
}

func (s *Service) SavePlane(ctx context.Context, p *plane.Plane) (*plane.Plane, error) {
	return s.db.SavePlane(ctx, p)
}

func (s *Service) GetPlaneById(ctx context.Context, id int) (*plane.Plane, error) {
	return s.db.GetPlaneById(ctx, id)
}

func (s *Service) UpdatePlane(ctx context.Context, p *plane.Plane) (*plane.Plane, error) {
	return s.db.UpdatePlane(ctx, p)
}

func (s *Service) GetAllPlanes(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[plane.Plane], error) {
	return s.db.GetAllPlanes(ctx, page, pageSize)
}

func (s *Service) GetPlanesBetweenDates(ctx context.Context, departureDate, arrivalDate time.Time, page, pageSize int) (database.PaginatedRecord[plane.Plane], error) {
	return s.db.GetPlanesBetweenDates(ctx, departureDate, arrivalDate, page, pageSize)
}
