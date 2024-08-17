package auth

import (
	"context"
	"fmt"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
)

// AuthService defines the interface for authentication-related operations
type AuthService interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email, password string) (string, error)
	Logout(ctx context.Context, userID int64) error
	ResetPassword(ctx context.Context, userID int64, newPassword string) error
	ShowAllUsers(ctx context.Context) ([]models.User, error)
}

type authService struct {
	AuthRepo AuthRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(AuthRepo AuthRepository) AuthService {
	return &authService{AuthRepo}
}

// ShowAllUsers retrieves all users from the repository
func (a *authService) ShowAllUsers(ctx context.Context) ([]models.User, error) {
	return a.AuthRepo.ShowAllUsers(ctx)
}

// Logout clears the user's token, effectively logging them out
func (a *authService) Logout(ctx context.Context, userID int64) error {
	return a.AuthRepo.UpdateUserToken(ctx, userID, "")
}

// ResetPassword updates the user's password after hashing it
func (a *authService) ResetPassword(ctx context.Context, userID int64, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	return a.AuthRepo.UpdatePassword(ctx, userID, hashedPassword)
}

// Login authenticates the user and returns a JWT token if successful
func (a *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.AuthRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	return token, nil
}

// FindUserByEmail retrieves a user by email from the repository
func (a *authService) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return a.AuthRepo.FindUserByEmail(ctx, email)
}

// CreateUser creates a new user after checking if the email already exists
func (a *authService) CreateUser(ctx context.Context, user *models.User) error {
	existingUser, err := a.AuthRepo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("failed to check existing user: %v", err)
	}
	if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user.Password = hashedPassword
	return a.AuthRepo.CreateUser(ctx, user)
}
