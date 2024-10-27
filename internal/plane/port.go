package plane

import (
	"context"
	"time"

	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Repository interface {
	SavePlane(ctx context.Context, p *Plane) (*Plane, error)
	GetPlaneById(ctx context.Context, id int) (*Plane, error)
	UpdatePlane(ctx context.Context, p *Plane) (*Plane, error)
	GetAllPlanes(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[Plane], error)
	GetPlanesBetweenDates(ctx context.Context, departureDate, arrivalDate time.Time, page, pageSize int) (database.PaginatedRecord[Plane], error)
}

