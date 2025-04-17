package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	UserCredentialRepository UserCredentialRepository
	UserProfileRepository    UserProfileRepository
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserProfileRepository:    NewUserProfileRepository(db),
		UserCredentialRepository: NewUserCredentialRepository(db),
	}
}
