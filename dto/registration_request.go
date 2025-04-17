package dto

type RegistrationRequest struct {
	ProfileName string `json:"profile_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
