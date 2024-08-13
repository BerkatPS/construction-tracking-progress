package project

import "net/http"

func RegisterRoutes(router *http.ServeMux, handler *ProjectController) {
	router.HandleFunc("GET /projects", handler.FindAll)
	//router.HandleFunc("GET /projects/{id}", handler.FindProjectByID)
	router.HandleFunc("POST /projects", handler.CreateProject)
	router.HandleFunc("PUT /projects/{id}", handler.UpdateProject)
	router.HandleFunc("DELETE /projects/{id}", handler.DeleteProject)
}
