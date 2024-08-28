package project

import "net/http"

func RegisterRoutes(router *http.ServeMux, handler *ProjectController) {
	router.HandleFunc("GET /projects", handler.FindAll)
	router.HandleFunc("GET /projects/{id}", handler.FindProjectByID)
	router.HandleFunc("POST /projects/add", handler.CreateProject)
	router.HandleFunc("PUT /projects/{id}", handler.UpdateProject)
	router.HandleFunc("DELETE /projects/{id}", handler.DeleteProject)
    router.HandleFunc("GET /projects/status-code/{status}", handler.FindProjectsByStatus) 
	router.HandleFunc("PUT /projects/{id}/status/{status}", handler.UpdateProjectStatus)
	router.HandleFunc("POST /projects/{id}/team/{user_id}", handler.AddTeamMemberToProject)
	router.HandleFunc("DELETE /projects/{id}/team/{user_id}", handler.RemoveTeamMemberFromProject)
	router.HandleFunc("PUT /projects/{id}/team/{user_id}/role/{role}", handler.UpdateProjectTeamRole)
	router.HandleFunc("POST /projects/{id}/expenses", handler.TrackProjectExpenses)
	router.HandleFunc("GET /projects/{id}/expenses", handler.FindExpensesByProject)
	router.HandleFunc("PUT /projects/{id}/budget/{new_budget}", handler.UpdateProjectBudget)
	router.HandleFunc("DELETE /projects/{id}/documents/{document_id}", handler.DeleteProjectDocument)
	router.HandleFunc("POST /projects/{id}/documents", handler.UploadProjectDocument)
}
