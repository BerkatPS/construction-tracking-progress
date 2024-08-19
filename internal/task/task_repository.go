package task

import (
	"context"
	"database/sql"
	models "github.com/BerkatPS/internal"
)

type TaskRepository interface {
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

}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) FindOverdueTasks(ctx context.Context) ([]models.Task, error) {
	query := "SELECT * FROM tasks WHERE end_date < CURRENT_DATE AND status = 'IN_PROGRESS'"

	rows, err := t.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.ProjectID, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) FindTasksByAssignedUser(ctx context.Context, userID int64) ([]models.Task, error) {
	query := "SELECT * FROM tasks WHERE assigned_user_id = $1"

	rows, err := t.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.ProjectID, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) TaskMarkAsInProgress(ctx context.Context, id int64) error {
	query := "UPDATE tasks SET status = 'IN_PROGRESS' WHERE id = $1"
	_, err := t.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}



func (t *taskRepository) TaskMarkAsDone(ctx context.Context, id int64) error {
	query := "UPDATE tasks SET status = 'DONE' WHERE id = $1"
	_, err := t.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) ArchiveCompletedTasks(ctx context.Context) error {
	query := "UPDATE tasks SET status = 'ARCHIVED' WHERE status = 'DONE'"
	_, err := t.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
}



func (t *taskRepository) ShowAllTasks(ctx context.Context) ([]models.Task, error) {
	query := "SELECT * FROM tasks"

	rows, err := t.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.ProjectID, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) FindTaskByID(ctx context.Context, id int64) (*models.Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1"

	row := t.db.QueryRowContext(ctx, query, id)
	var task models.Task
	if err := row.Scan(&task.ID, &task.ProjectID, &task.Name, &task.Description, &task.StartDate, &task.EndDate, &task.Status); err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *taskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	query := "INSERT INTO tasks (project_id, name, description, start_date, end_date, status) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := t.db.ExecContext(ctx, query, task.ProjectID, task.Name, task.Description, task.StartDate, task.EndDate, task.Status)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) UpdateTask(ctx context.Context, task *models.Task) error {
	query := "UPDATE tasks SET project_id = $1, name = $2, description = $3, start_date = $4, end_date = $5, status = $6 WHERE id = $7"

	_, err := t.db.ExecContext(ctx, query, task.ProjectID, task.Name, task.Description, task.StartDate, task.EndDate, task.Status, task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) DeleteTask(ctx context.Context, id int64) error {
	query := "DELETE FROM tasks WHERE id = $1"

	_, err := t.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
