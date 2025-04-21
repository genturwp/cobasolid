package service

import (
	"cobasolid/dto"
	"cobasolid/entity"
	"cobasolid/repository"
	"context"
)

type UserManagementService interface {
	UserRegistration(ctx context.Context, registrationReq *dto.RegistrationRequest) (*entity.UserProfile, error)
	ChangePassword(ctx context.Context, oldPassword string, newPassword string) error
}

type userManagementSvc struct {
	repo *repository.Repository
}

func NewUserManagementService(repository *repository.Repository) UserManagementService {
	return &userManagementSvc{
		repo: repository,
	}
}

func (svc *userManagementSvc) UserRegistration(ctx context.Context, registrationReq *dto.RegistrationRequest) (*entity.UserProfile, error) {
	return nil, nil
}

func (svc *userManagementSvc) ChangePassword(ctx context.Context, oldPassword string, newPassword string) error {
	return nil
}
