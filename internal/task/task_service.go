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
}

type taskService struct {
	TaskRepo TaskRepository
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

func NewTaskService(taskRepo TaskRepository) TaskService {
	return &taskService{taskRepo}
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
