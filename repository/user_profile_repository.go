package repository

import (
	"cobasolid/entity"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserProfileRepository interface {
	Save(ctx context.Context, userProfile *entity.UserProfile) (*entity.UserProfile, error)
	Update(ctx context.Context, userProfile *entity.UserProfile) (*entity.UserProfile, error)
	FindById(ctx context.Context, id uuid.UUID) (*entity.UserProfile, error)
	FindAll(ctx context.Context) ([]entity.UserProfile, error)
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
	q := `
		insert into user_profiles(profile_name, phone, email)
		values(@profileName, @phone, @email)
		returning *
	`
	args := pgx.NamedArgs{
		"profileName": userProfile.ProfileName,
		"phone":       userProfile.Phone,
		"email":       userProfile.Email,
	}
	rows, err := repo.DB.Query(ctx, q, args)
	if err != nil {
		return nil, err
	}
	return pgx.CollectExactlyOneRow(rows, pgx.RowToAddrOfStructByName[entity.UserProfile])
}

func (repo *userProfileRepo) Update(ctx context.Context, userProfile *entity.UserProfile) (*entity.UserProfile, error) {
	return nil, nil
}
func (repo *userProfileRepo) FindById(ctx context.Context, id uuid.UUID) (*entity.UserProfile, error) {
	return nil, nil
}
func (repo *userProfileRepo) FindAll(ctx context.Context) ([]entity.UserProfile, error) {
	return nil, nil
}
