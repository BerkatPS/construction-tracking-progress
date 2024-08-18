package server

import (
	"database/sql"
	"github.com/BerkatPS/internal/auth"
	"github.com/BerkatPS/internal/project"
	"github.com/BerkatPS/internal/quality"
	"github.com/BerkatPS/internal/task"
	"github.com/BerkatPS/pkg/middleware"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
	db     *sql.DB
}

func NewServer(db *sql.DB) *Server {
	router := http.NewServeMux()
	s := &Server{
		Router: router,
		db:     db,
	}

	//router.Use(middleware.LoggingMiddleware)
	//router.Use(middleware.RecoveryMiddleware)
	//router.Use(middleware.CORSHandler)
	//router.Use(middleware.AuthMiddleware)

	s.registerRoutes()

	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	return s
}

func (s *Server) registerRoutes() {
	// auth routes
	authRepo := auth.NewAuthRepository(s.db)
	authService := auth.NewAuthService(authRepo)
	authController := auth.NewAuthController(authService)
	auth.RegisterRoutes(s.Router, authController)

	// project routes
	projectRepo := project.NewProjectRepository(s.db)
	projectService := project.NewProjectService(projectRepo)
	projectController := project.NewProjectController(projectService)
	project.RegisterRoutes(s.Router, projectController)

	// expenses routes

	// Document routes

	// Project routes

	// QualityCheck routes

	// SafetyIncident routes

	// report routes

	// Task Routes
	taskRepo := task.NewTaskRepository(s.db)
	taskService := task.NewTaskService(taskRepo)
	taskController := task.NewTaskController(taskService)
	task.RegisterRoutes(s.Router, taskController)

	// Message Routes

	// quality Routes
	qualityRepo := quality.NewQualityRepository(s.db)
	qualityService := quality.NewQualityService(qualityRepo)
	qualityController := quality.NewQualityController(qualityService)
	quality.RegisterRoutes(s.Router, qualityController)

}

func (s *Server) applyMiddleware(h http.Handler) http.Handler {
	return middleware.LoggingMiddleware(
		middleware.RecoveryMiddleware(
			middleware.CORSHandler(
				middleware.AuthMiddleware(h),
			),
		),
	)
}

// exampleHandler is an example of a simple route handler
func (s *Server) exampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

// loggingMiddleware logs each request
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logging logic here
		next.ServeHTTP(w, r)
	})
}

// recoveryMiddleware recovers from panics
func (s *Server) recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware adds CORS headers to the response
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS logic here
		next.ServeHTTP(w, r)
	})
}

// authMiddleware checks for a valid authentication token
func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authentication logic here
		next.ServeHTTP(w, r)
	})
}
