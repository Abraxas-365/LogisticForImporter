package consignatario

import (
	"context"
	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Repository interface {
	SaveConsignatario(ctx context.Context, u *Consignatario) (*Consignatario, error)
	GetConsignatarioById(ctx context.Context, id int) (*Consignatario, error)
	UpdateConsignatario(ctx context.Context, u *Consignatario) (*Consignatario, error)
	GetAllByUserId(ctx context.Context, userId int, page int, pageSize int) (database.PaginatedRecord[Consignatario], error)
}
