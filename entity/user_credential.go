package entity

import "github.com/google/uuid"

type UserCredential struct {
	BaseEntity
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	IsEnabled bool       `json:"is_enabled"`
	Roles     []string   `json:"roles"`
	ProfileId *uuid.UUID `json:"profile_id"`
}
