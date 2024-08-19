package task

import (
	"context"
	"fmt"
	models "github.com/BerkatPS/internal"
)

type TaskService interface {
	ShowAllTasks(ctx context.Context) ([]models.Task, error)
	FindTaskByID(ctx context.Context, id int64) (*models.Task, error)
	CreateTask(ctx context.Context, task *models.Task) error
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, id int64) error
	TaskMarkAsDone(ctx context.Context, id int64) error
	ArchiveCompletedTasks(ctx context.Context) error
	TaskMarkAsInProgress(ctx context.Context, id int64) error
	FindTasksByAssignedUser(ctx context.Context, userID int64) ([]models.Task, error)
	FindOverdueTasks(ctx context.Context) ([]models.Task, error)
	FindTasksByProjectID(ctx context.Context, projectID int64) ([]models.Task, error)
}

type taskService struct {
	TaskRepo TaskRepository
}


func NewTaskService(taskRepo TaskRepository) TaskService {
	return &taskService{taskRepo}
}

func (t *taskService) FindTasksByProjectID(ctx context.Context, projectID int64) ([]models.Task, error) {
	tasks, err := t.TaskRepo.FindTasksByProjectID(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks for project: %v", err)
	}

	return tasks, nil
}

func (t *taskService) FindOverdueTasks(ctx context.Context) ([]models.Task, error) {
	tasks, err := t.TaskRepo.FindOverdueTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve overdue tasks: %v", err)
	}

	return tasks, nil
}

func (t *taskService) FindTasksByAssignedUser(ctx context.Context, userID int64) ([]models.Task, error){
	tasks, err := t.TaskRepo.FindTasksByAssignedUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks for user: %v", err)
	}

	return tasks, nil
}

func (t *taskService) ArchiveCompletedTasks(ctx context.Context) error {
	err := t.TaskRepo.ArchiveCompletedTasks(ctx)
	if err != nil {
		return fmt.Errorf("failed to archive completed tasks: %v", err)
	}

	return nil
}

func (t *taskService) TaskMarkAsInProgress(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid task ID")
	}
	err := t.TaskRepo.TaskMarkAsInProgress(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to mark task as in progress: %v", err)
	}

	return nil
}
func (t *taskService) TaskMarkAsDone(ctx context.Context, id int64) error {

	if id <= 0 {
		return fmt.Errorf("invalid task ID")
	}
	err := t.TaskRepo.TaskMarkAsDone(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to mark task as done: %v", err)
	}

	return nil
}


func (t *taskService) ShowAllTasks(ctx context.Context) ([]models.Task, error) {
	tasks, err := t.TaskRepo.ShowAllTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %v", err)
	}

	return tasks, nil
}

func (t *taskService) FindTaskByID(ctx context.Context, id int64) (*models.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid task ID")
	}

	task, err := t.TaskRepo.FindTaskByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve task: %v", err)
	}

	return task, nil
}

func (t *taskService) CreateTask(ctx context.Context, task *models.Task) error {
	if task.Name == "" || task.Description == "" {
		return fmt.Errorf("missing required task fields")
	}

	if err := t.TaskRepo.CreateTask(ctx, task); err != nil {
		return fmt.Errorf("failed to create task: %v", err)
	}

	return nil
}

func (t *taskService) UpdateTask(ctx context.Context, task *models.Task) error {
	if task.ID <= 0 || task.Name == "" || task.Description == "" {
		return fmt.Errorf("missing required task fields")
	}

	err := t.TaskRepo.UpdateTask(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	return nil
}

func (t *taskService) DeleteTask(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid task ID")
	}

	err := t.TaskRepo.DeleteTask(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %v", err)
	}

	return nil
}
