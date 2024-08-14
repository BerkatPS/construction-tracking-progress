package auth

import (
	"fmt"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
)

type AuthService interface {
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	Login(email, password string) error
	Logout(userID int64) error
	//RefreshToken(refreshToken string) (string, error)
	ResetPassword(userID int64, newPassword string) error
	ShowAllUsers() ([]models.User, error)
}

type authService struct {
	Authrepo AuthRepository
}

func NewAuthService(Authrepo AuthRepository) AuthService {
	return &authService{Authrepo}
}

func (a *authService) ShowAllUsers() ([]models.User, error) {
	return a.Authrepo.ShowAllUsers()
}

func (a *authService) Logout(userID int64) error {
	return a.Authrepo.UpdateUserToken(userID, "")
}

//func (a *authService) RefreshToken(refreshToken string) (string, error) {
//	user, err := a.Authrepo.FindUserByID()
//	if err != nil {
//		return "", err
//	}
//	if user.RefreshToken != refreshToken {
//		return "", fmt.Errorf("invalid refresh token")
//	}
//	//generate jwt token
//	newtoken, err := utils.GenerateToken(int(user.ID))
//	if err != nil {
//		return "", fmt.Errorf("failed to generate Token: %v", err)
//	}
//	return newtoken, nil
//}

func (a *authService) ResetPassword(userID int64, newPassword string) error {

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	return a.Authrepo.UpdatePassword(userID, hashedPassword)
}

func (a *authService) Login(email, password string) error {
	user, err := a.Authrepo.FindUserByEmail(email)
	if err != nil {
		return err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return fmt.Errorf("invalid password")
	}

	//generate jwt token
	_, err = utils.GenerateToken(int(user.ID))
	if err != nil {
		return fmt.Errorf("failed to generate Token: %v", err)
	}
	return nil
}

func (a *authService) FindUserByEmail(email string) (*models.User, error) {
	user, err := a.Authrepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a authService) CreateUser(user *models.User) error {

	email, err := a.Authrepo.FindUserByEmail(user.Email)
	if email != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)

	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v")
	}

	user.Password = hashedPassword
	return a.Authrepo.CreateUser(user)
}
