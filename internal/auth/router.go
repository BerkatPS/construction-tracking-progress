package auth

import "net/http"

func RegisterRoutes(router *http.ServeMux, handler *AuthController) {
	router.HandleFunc("POST /register", handler.CreateUser)
	router.HandleFunc("POST /login", handler.Login)
	router.HandleFunc("POST /logout", handler.Logout)
	//router.HandleFunc("POST /refresh", handler.RefreshToken)
	router.HandleFunc("POST /reset-password", handler.ResetPassword)
}
