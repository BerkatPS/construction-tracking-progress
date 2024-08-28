package middleware

import (
	"github.com/BerkatPS/pkg/utils"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))

		next.ServeHTTP(w, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v\n%s", err, debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		authHeader := request.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(writer, "Forbidden: no token provided", http.StatusForbidden)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !utils.ValidateToken(token) {
			http.Error(writer, "Forbidden: invalid token provided", http.StatusForbidden)
			return
		}

		next.ServeHTTP(writer, request)

	})
}

func CORSHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}

func IPMiddleware(allowedIPs []string) func(http.Handler) http.Handler {
	return func (next http.Handler) http.Handler  {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.Header.Get("X-Forwarded-For")
			if ip == "" {
				ip = r.RemoteAddr
			}
			r.Header.Set("X-Forwarded-For", ip)
			next.ServeHTTP(w, r)
		})
	}
}


func isAllowedIP(ip string, allowedIPs []string) bool {
	for _, allowedIP := range allowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}