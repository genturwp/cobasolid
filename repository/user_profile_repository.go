package repository

import (
	"cobasolid/entity"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserProfileRepository interface {
	Save(ctx context.Context, userProfile *entity.UserProfile) (*entity.UserProfile, error)
	Update(ctx context.Context, userProfile *entity.UserProfile) (*entity.UserProfile, error)
	FindById(ctx context.Context, id uuid.UUID) (*entity.UserProfile, error)
}

type userProfileRepo struct {
	DB *pgxpool.Pool
}

func NewUserProfileRepository(db *pgxpool.Pool) UserProfileRepository {
	return &userProfileRepo{
		DB: db,
	}
}
func (repo *userProfileRepo) Save(ctx context.Context, userProfile *entity.UserProfile) (*entity.UserProfile, error) {
	return nil, nil
}
func (repo *userProfileRepo) Update(ctx context.Context, userProfile *entity.UserProfile) (*entity.UserProfile, error) {
	return nil, nil
}
func (repo *userProfileRepo) FindById(ctx context.Context, id uuid.UUID) (*entity.UserProfile, error) {
	return nil, nil
}
