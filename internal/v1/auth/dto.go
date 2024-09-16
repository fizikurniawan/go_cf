package auth

type RegisterUserDTO struct {
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RegisterBy  string `json:"register_by" binding:"required"`
}
