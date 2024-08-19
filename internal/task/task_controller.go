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

func (t *TaskController) FindTasksByProjectID (w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	projectID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	tasks, err := t.Service.FindTasksByProjectID(ctx, projectID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve tasks for project: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Tasks found successfully",
		"data":    tasks,
	})
}

func (t *TaskController) TaskMarkAsInProgress(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid task ID: " + err.Error(),
		})
		return
	}

	err = t.Service.TaskMarkAsInProgress(ctx, taskID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to mark task as in progress: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Task marked as in progress successfully",
	})
}

func (t *TaskController) FindOverdueTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tasks, err := t.Service.FindOverdueTasks(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve overdue tasks: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Overdue tasks found successfully",
		"data":    tasks,
	})
}

func (t *TaskController) FindTasksByAssignedUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid user ID: " + err.Error(),
		})
		return
	}

	tasks, err := t.Service.FindTasksByAssignedUser(ctx, userID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve tasks for user: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Tasks found successfully",
		"data":    tasks,
	})
}

func (t *TaskController) ArchiveCompletedTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := t.Service.ArchiveCompletedTasks(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to archive completed tasks: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Completed tasks archived successfully",
	})
}

func (t *TaskController) TaskMarkAsDone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	taskID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid task ID: " + err.Error(),
		})
		return
	}

	err = t.Service.TaskMarkAsDone(ctx, taskID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to mark task as done: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Task marked as done successfully",
	})
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
