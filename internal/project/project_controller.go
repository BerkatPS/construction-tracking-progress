package project

import (
	"encoding/json"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
	"net/http"
)

type ProjectController struct {
	projectService ProjectService
}

func NewProjectController(projectService ProjectService) *ProjectController {
	return &ProjectController{projectService}
}

func (pc *ProjectController) FindProjectsByStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	status := r.URL.Query().Get("status")
	if status == "" {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Missing required query parameter 'status'",
		})
		return
	}

	projects, err := pc.projectService.FindProjectsByStatus(ctx, status)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve projects: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Projects found successfully",
		"data":    projects,
	})
}


func (pc *ProjectController) UpdateProjectStatus(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	var updateProjectStatusRequest struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateProjectStatusRequest); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.UpdateProjectStatus(ctx, id, updateProjectStatusRequest.Status); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update project status: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project status updated successfully",
	})
}

func (pc *ProjectController) AddTeamMemberToProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	var addTeamMemberToProjectRequest struct {
		UserID int64 `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&addTeamMemberToProjectRequest); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.AddTeamMemberToProject(ctx, id, addTeamMemberToProjectRequest.UserID); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to add team member to project: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Team member added to project successfully",
	})
}

func (pc *ProjectController) RemoveTeamMemberFromProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	var removeTeamMemberFromProjectRequest struct {
		UserID int64 `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&removeTeamMemberFromProjectRequest); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.RemoveTeamMemberFromProject(ctx, id, removeTeamMemberFromProjectRequest.UserID); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to remove team member from project: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Team member removed from project successfully",
	})
}

func (pc *ProjectController) UpdateProjectTeamRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	var updateProjectTeamRoleRequest struct {
		UserID int64  `json:"user_id"`
		Role   string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateProjectTeamRoleRequest); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.UpdateProjectTeamRole(ctx, id, updateProjectTeamRoleRequest.UserID, updateProjectTeamRoleRequest.Role); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update project team role: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project team role updated successfully",
	})
}

func (pc *ProjectController) TrackProjectExpenses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	var trackProjectExpensesRequest struct {
		Amount      float64 `json:"amount"`
		Description string  `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&trackProjectExpensesRequest); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	expense := &models.Expense{
		ProjectID:   id,
		Amount:      trackProjectExpensesRequest.Amount,
		Description: trackProjectExpensesRequest.Description,
	}

	if err := pc.projectService.TrackProjectExpenses(ctx, expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to track project expenses: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project expenses tracked successfully",
	})
}

func (pc *ProjectController) FindExpensesByProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	expenses, err := pc.projectService.FindExpensesByProject(ctx, id)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve expenses: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Expenses found successfully",
		"data":    expenses,
	})
}


func (pc *ProjectController) UpdateProjectBudget(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	var updateProjectBudgetRequest struct {
		NewBudget float64 `json:"new_budget"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateProjectBudgetRequest); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.UpdateProjectBudget(ctx, id, updateProjectBudgetRequest.NewBudget); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update project budget: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project budget updated successfully",
	})
}

func (pc *ProjectController) DeleteProjectDocument(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.DeleteProjectDocument(ctx, id); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to delete project document: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project document deleted successfully",
	})
}

func (pc *ProjectController) UploadProjectDocument(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	var uploadProjectDocumentRequest struct {
		Document *models.Document `json:"document"`
	}

	if err := json.NewDecoder(r.Body).Decode(&uploadProjectDocumentRequest); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.UploadProjectDocument(ctx, id, uploadProjectDocumentRequest.Document); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to upload project document: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project document uploaded successfully",
	})
}

// FindAll handles the request to retrieve all projects
func (pc *ProjectController) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projects, err := pc.projectService.FindAll(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve projects: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Projects found successfully",
		"data":    projects,
	})
}

// FindProjectByID handles the request to retrieve a project by its ID
func (pc *ProjectController) FindProjectByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	project, err := pc.projectService.FindProjectByID(ctx, id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve project: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project found successfully",
		"data":    project,
	})
}

// CreateProject handles the request to create a new project
func (pc *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.CreateProject(ctx, &project); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create project: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project created successfully",
	})
}

// UpdateProject handles the request to update an existing project
func (pc *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.UpdateProject(ctx, &project); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update project: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project updated successfully",
	})
}

// DeleteProject handles the request to delete a project by its ID
func (pc *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	if err := pc.projectService.DeleteProject(ctx, id); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to delete project: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Project deleted successfully",
	})
}
