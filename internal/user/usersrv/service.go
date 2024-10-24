package usersrv

import (
	"context"
	"log"

	"github.com/Abraxas-365/cabo/internal/consignee"
	"github.com/Abraxas-365/cabo/internal/user"
	"github.com/Abraxas-365/cabo/pkg/apis"
	"github.com/Abraxas-365/toolkit/pkg/database"
	"github.com/Abraxas-365/toolkit/pkg/errors"
)

type Service struct {
	repo             user.Repository
	consignatarioSrv apis.ConsigneeServicer
}

func New(repo user.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, u *user.User) error {
	newUser, err := s.repo.SaveUser(ctx, u)
	if err != nil {
		return err
	}

	docType, err := consignee.NewDocumentType(string(newUser.DocumentType))
	if err != nil {
		return err
	}

	consignatario := &consignee.Consignee{
		UserID:         newUser.ID,
		Phone:          newUser.Phone,
		Email:          newUser.Email,
		DocumentType:   docType,
		DocumentNumber: newUser.DocumentNumber,
	}

	_, err = s.consignatarioSrv.CreateConsignatario(ctx, consignatario)
	if err != nil {
		//TODO: add this to a queue to try crating the consignatario again
		log.Println("Error creating consignatario", err)
	}

	return nil
}

func (s *Service) GetUser(ctx context.Context, id int) (*user.User, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, u *user.User) (*user.User, error) {
	return s.repo.UpdateUser(ctx, u)
}

func (s *Service) GetAllUsers(ctx context.Context, page int, pageSize int) (database.PaginatedRecord[user.User], error) {
	return s.repo.GetAllUsers(ctx, page, pageSize)
}

func (s *Service) SaveDirection(ctx context.Context, d *user.Direction) (*user.Direction, error) {
	return s.repo.SaveDirection(ctx, d)
}

func (s *Service) UpdateDirection(ctx context.Context, d *user.Direction) (*user.Direction, error) {
	if ad, err := s.repo.GetDirectionByID(ctx, d.ID); err != nil {
		return nil, err
	} else if ad.UserID != d.UserID {
		return nil, errors.ErrForbidden("User not allowed to access this resource")
	}

	return s.repo.UpdateDirection(ctx, d)
}

func (s *Service) GetUserDirections(ctx context.Context, userID int, page int, pageSize int) (database.PaginatedRecord[user.Direction], error) {
	return s.repo.GetUserDirections(ctx, userID, page, pageSize)
}
