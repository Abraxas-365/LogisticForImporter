package user

import (
	"context"

	"github.com/Abraxas-365/toolkit/pkg/database"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, u *User) error {
	return s.repo.SaveUser(ctx, u)
}

func (s *Service) GetUser(ctx context.Context, id int) (*User, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, u *User) (*User, error) {
	return s.repo.UpdateUser(ctx, u)
}

func (s *Service) GetAllUsers(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[User], error) {
	return s.repo.GetAllUsers(ctx, page, pageSize)
}
