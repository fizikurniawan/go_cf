// internal/v1/auth/service.go
package auth

import (
	"crowdfunding/internal/common/models"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	DB *sqlx.DB
}

// NewService creates a new auth service
func NewService(db *sqlx.DB) *Service {
	return &Service{DB: db}
}

// RegisterUser handles user registration
func (s *Service) RegisterUser(dto RegisterUserDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		ID:          uuid.New(),
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Password:    string(hashedPassword),
		IsActive:    true,
		Role:        "user", // Default role
		RegisterBy:  dto.RegisterBy,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	query := `INSERT INTO users (id, email, phone_number, first_name, last_name, password, is_active, role, register_by, created_at, updated_at) 
              VALUES (:id, :email, :phone_number, :first_name, :last_name, :password, :is_active, :role, :register_by, :created_at, :updated_at)`

	_, err = s.DB.NamedExec(query, user)
	return err
}

func (s *Service) GetByEmail(email string) (models.User, error) {
	var user models.User

	// Query dengan parameterized SQL untuk PostgreSQL
	query := "SELECT * FROM users WHERE email = $1"

	// Eksekusi query dengan parameter
	err := s.DB.Get(&user, query, email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
