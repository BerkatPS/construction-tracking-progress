package auth

import (
	"database/sql"
	"errors"
	models "github.com/BerkatPS/internal"
)

type AuthRepository interface {
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUserToken(userID int64, token string) error      // Untuk menyimpan token refresh
	FindUserByID(userID int64) (*models.User, error)       // Untuk mencari user berdasarkan ID
	UpdatePassword(userID int64, newPassword string) error // Untuk mengupdate password user

}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) UpdatePassword(userID int64, newPassword string) error {
	query := "UPDATE users SET password = $1 WHERE id = $2"
	_, err := r.db.Exec(query, newPassword, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) FindUserByID(userID int64) (*models.User, error) {
	query := "SELECT id, username, email, password, role FROM users WHERE id = $1"

	row := r.db.QueryRow(query, userID)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *authRepository) UpdateUserToken(userID int64, token string) error {
	query := "UPDATE users SET refresh_token = $1 WHERE id = $2"
	_, err := r.db.Exec(query, token, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) FindUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password, role FROM users WHERE email = $1"

	row := r.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func (r *authRepository) CreateUser(user *models.User) error {

	query := "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4)"

	_, err := r.db.Exec(query, user.Username, user.Email, user.Password, user.Role)

	if err != nil {
		return err
	}
	return nil
}
