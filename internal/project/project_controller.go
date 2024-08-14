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
