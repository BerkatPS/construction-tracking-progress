package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	models "github.com/BerkatPS/internal"
)

const (
	selectUserByIDQuery    = "SELECT id, username, email, password, role FROM users WHERE id = $1"
	selectUserByEmailQuery = "SELECT id, username, email, password, role FROM users WHERE email = $1"
	insertUserQuery        = "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)"
	updatePasswordQuery    = "UPDATE users SET password = $1 WHERE id = $2"
	updateUserTokenQuery   = "UPDATE users SET refresh_token = $1 WHERE id = $2"
	selectAllUsersQuery    = "SELECT id, username, email, password, role FROM users"
)

// AuthRepository defines the methods for interacting with user data
type AuthRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUserToken(ctx context.Context, userID int64, token string) error
	FindUserByID(ctx context.Context, userID int64) (*models.User, error)
	UpdatePassword(ctx context.Context, userID int64, newPassword string) error
	ShowAllUsers(ctx context.Context) ([]models.User, error)
}

type authRepository struct {
	db *sql.DB
}

// NewAuthRepository creates a new instance of AuthRepository
func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db}
}

// ShowAllUsers retrieves all users from the database
func (r *authRepository) ShowAllUsers(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.QueryContext(ctx, selectAllUsersQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %w", err)
	}
	return users, nil
}

// UpdatePassword updates the password for a user
func (r *authRepository) UpdatePassword(ctx context.Context, userID int64, newPassword string) error {
	_, err := r.db.ExecContext(ctx, updatePasswordQuery, newPassword, userID)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}
	return nil
}

// FindUserByID retrieves a user by their ID
func (r *authRepository) FindUserByID(ctx context.Context, userID int64) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, selectUserByIDQuery, userID)

	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}
	return user, nil
}

// UpdateUserToken updates the refresh token for a user
func (r *authRepository) UpdateUserToken(ctx context.Context, userID int64, token string) error {
	_, err := r.db.ExecContext(ctx, updateUserTokenQuery, token, userID)
	if err != nil {
		return fmt.Errorf("failed to update user token: %w", err)
	}
	return nil
}

// FindUserByEmail retrieves a user by their email
func (r *authRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := r.db.QueryRowContext(ctx, selectUserByEmailQuery, email)

	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	return user, nil
}

// CreateUser adds a new user to the database
func (r *authRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, insertUserQuery, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}
