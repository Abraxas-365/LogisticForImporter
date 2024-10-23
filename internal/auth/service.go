package auth

import (
	"context"

	"github.com/Abraxas-365/toolkit/pkg/lucia"
)

type Service struct{}

func (s *Service) GetUserByProviderID(ctx context.Context, provider, providerID string) (*AuthUser, error) {
	return nil, nil
}

func (s *Service) CreateUser(ctx context.Context, userInfo *lucia.UserInfo) (*AuthUser, error) {
	return nil, nil
}
