package apis

import (
	"context"
	"os/user"

	"github.com/Abraxas-365/toolkit/pkg/database"
)

type UserServicer interface {
	CreateUser(ctx context.Context, u *user.User) error

	GetUser(ctx context.Context, id int) (*user.User, error)

	UpdateUser(ctx context.Context, u *user.User) (*user.User, error)

	GetAllUsers(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[user.User], error)
}
