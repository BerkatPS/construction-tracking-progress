package task

import "net/http"

func RegisterRoutes(router *http.ServeMux, handler *TaskController) {
	router.HandleFunc("GET /tasks", handler.ShowAllTasks)
	router.HandleFunc("GET /tasks/{id}", handler.FindTaskByID)
	router.HandleFunc("POST /tasks/add", handler.CreateTask)
	router.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	router.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)
	router.HandleFunc("PUT /tasks/{id}/done", handler.TaskMarkAsDone)
	router.HandleFunc("PUT /tasks/{id}/in-progress", handler.TaskMarkAsInProgress)
	router.HandleFunc("GET /tasks/overdue", handler.FindOverdueTasks)
	router.HandleFunc("GET /tasks/user/{user_id}", handler.FindTasksByAssignedUser)
	router.HandleFunc("PUT /tasks/archive", handler.ArchiveCompletedTasks)
}
