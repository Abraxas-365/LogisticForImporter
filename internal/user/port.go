package user

import (
	"context"

	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Repository interface {
	SaveUser(ctx context.Context, u *User) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	GetAllUsers(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[User], error)
	SaveDirection(ctx context.Context, d *Direction) (*Direction, error)
	UpdateDirection(ctx context.Context, d *Direction) (*Direction, error)
	GetUserDirections(ctx context.Context, userID int, page int, pageSize int) (database.PaginatedRecord[Direction], error)
	GetDirectionByID(ctx context.Context, id int) (*Direction, error)
}
