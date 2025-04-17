package service

import "cobasolid/repository"

type Service struct {
	UserManagementService UserManagementService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserManagementService: NewUserManagementService(repo),
	}
}
