package apis

import (
	"context"

	"github.com/Abraxas-365/cabo/internal/consignee"
	"github.com/Abraxas-365/toolkit/pkg/database"
)

type ConsigneeServicer interface {
	CreateConsignatario(ctx context.Context, consignatario *consignee.Consignee) (*consignee.Consignee, error)

	GetConsignatario(ctx context.Context, id int) (*consignee.Consignee, error)

	UpdateConsignatario(ctx context.Context, consignatario *consignee.Consignee) (*consignee.Consignee, error)

	GetAllByUserId(ctx context.Context, userId int, pageNumber, pageSize int) (database.PaginatedRecord[consignee.Consignee], error)
}
