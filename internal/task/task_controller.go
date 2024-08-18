package task

import (
	"encoding/json"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
	"net/http"
)

type TaskController struct {
	Service TaskService
}

func NewTaskController(service TaskService) *TaskController {
	return &TaskController{service}
}

func (t *TaskController) ShowAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tasks, err := t.Service.ShowAllTasks(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve tasks: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Tasks found successfully",
		"data":    tasks,
	})
}

func (t *TaskController) FindTaskByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid task ID: " + err.Error(),
		})
		return
	}

	task, err := t.Service.FindTaskByID(ctx, id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve task: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Task found successfully",
		"data":    task,
	})
}

func (t *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := t.Service.CreateTask(ctx, &task); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create task: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Task created successfully",
	})
}

func (t *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request payload: " + err.Error(),
		})
		return
	}

	if err := t.Service.UpdateTask(ctx, &task); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update task: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Task updated successfully",
	})
}

func (t *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid task ID: " + err.Error(),
		})
		return
	}

	if err := t.Service.DeleteTask(ctx, id); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to delete task: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Task deleted successfully",
	})
}
