package auth

import (
	"encoding/json"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
	"net/http"
)

// AuthController handles HTTP requests related to authentication
type AuthController struct {
	AuthService AuthService
}

// NewAuthController creates a new AuthController instance
func NewAuthController(authService AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// ShowAllUsers retrieves all users and returns them in the response
func (a *AuthController) ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the context from the request

	users, err := a.AuthService.ShowAllUsers(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve users: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

// CreateUser creates a new user from the provided request data
func (a *AuthController) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the context from the request

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := a.AuthService.CreateUser(ctx, &user); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create user: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User created successfully",
	})
}

// Login authenticates a user and returns a JWT token if successful
func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the context from the request

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid credentials payload: " + err.Error(),
		})
		return
	}

	token, err := a.AuthService.Login(ctx, credentials.Email, credentials.Password)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusUnauthorized, map[string]interface{}{
			"status":  "error",
			"message": "Authentication failed: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Login successful",
		"token":   token,
	})
}

// Logout invalidates the user's session by clearing their token
func (a *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the context from the request

	userID, ok := r.Context().Value("userID").(int64)
	if !ok {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "User ID not found in context",
		})
		return
	}

	if err := a.AuthService.Logout(ctx, userID); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to logout user: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User logged out successfully",
	})
}

// ResetPassword updates the user's password
func (a *AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Get the context from the request

	var request struct {
		UserID   int64  `json:"user_id"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid reset password payload: " + err.Error(),
		})
		return
	}

	if err := a.AuthService.ResetPassword(ctx, request.UserID, request.Password); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to reset password: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Password reset successfully",
	})
}
