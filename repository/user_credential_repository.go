package repository

import (
	"cobasolid/entity"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserCredentialRepository interface {
	Save(ctx context.Context, userCredential *entity.UserCredential) (*entity.UserCredential, error)
	Update(ctx context.Context, userCredential *entity.UserCredential) (*entity.UserCredential, error)
	FindByUsername(ctx context.Context, username string) (*entity.UserCredential, error)
}

type userCredentialRepo struct {
	DB *pgxpool.Pool
}

func NewUserCredentialRepository(db *pgxpool.Pool) UserCredentialRepository {
	return &userCredentialRepo{
		DB: db,
	}
}

func (repo *userCredentialRepo) Save(ctx context.Context, userCredential *entity.UserCredential) (*entity.UserCredential, error) {
	return nil, nil
}
func (repo *userCredentialRepo) Update(ctx context.Context, userCredential *entity.UserCredential) (*entity.UserCredential, error) {
	return nil, nil
}
func (repo *userCredentialRepo) FindByUsername(ctx context.Context, username string) (*entity.UserCredential, error) {
	return nil, nil
}
