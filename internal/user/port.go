package user

import (
	"context"

	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Repository interface {
	SaveUser(ctx context.Context, u *User) error
	GetUserById(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	GetAllUsers(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[User], error)
}
