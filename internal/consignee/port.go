package consignee

import (
	"context"
	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Repository interface {
	SaveConsignatario(ctx context.Context, u *Consignee) (*Consignee, error)
	GetConsignatarioById(ctx context.Context, id int) (*Consignee, error)
	UpdateConsignatario(ctx context.Context, u *Consignee) (*Consignee, error)
	GetAllByUserId(ctx context.Context, userId int, page int, pageSize int) (database.PaginatedRecord[Consignee], error)
}
