package auth

import (
	"encoding/json"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
	"net/http"
)

type AuthController struct {
	AuthService AuthService
}

func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (a *AuthController) ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.AuthService.ShowAllUsers()
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})

}

func (a *AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error()})
		return
	}

	if err := a.AuthService.CreateUser(&user); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User created successfully",
	})
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := a.AuthService.Login(credentials.Email, credentials.Password); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User logged in successfully",
	})
}

func (a *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	if err := a.AuthService.Logout(userID); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User logged out successfully",
	})
}

func (a *AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {

	var request struct {
		UserID   int64  `json:"user_id"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if err := a.AuthService.ResetPassword(request.UserID, request.Password); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Password reset successfully",
	})
}
