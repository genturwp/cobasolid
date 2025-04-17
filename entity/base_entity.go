package entity

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	Id        *uuid.UUID `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	CreatedBy *uuid.UUID `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *uuid.UUID `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *time.Time `json:"deleted_by"`
}
