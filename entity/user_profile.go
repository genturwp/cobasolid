package entity

type UserProfile struct {
	BaseEntity
	ProfileName string `json:"profile_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}
