package project

import (
	"context"
	"database/sql"
	"errors"
	models "github.com/BerkatPS/internal"
)

type ProjectRepository interface {
	FindAll(ctx context.Context) ([]models.Project, error)
	FindProjectByID(ctx context.Context, id int64) (*models.Project, error)
	CreateProject(ctx context.Context, project *models.Project) error
	UpdateProject(ctx context.Context, project *models.Project) error
	DeleteProject(ctx context.Context, id int64) error
}

type projectRepository struct {
	db *sql.DB
}

// NewProjectRepository creates a new instance of ProjectRepository
func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db}
}

// FindAll retrieves all projects from the database
func (p *projectRepository) FindAll(ctx context.Context) ([]models.Project, error) {
	query := "SELECT id, name, description, budget, status FROM projects"

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.Budget, &project.Status); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return projects, nil
}

// FindProjectByID retrieves a project by its ID from the database
func (p *projectRepository) FindProjectByID(ctx context.Context, id int64) (*models.Project, error) {
	query := "SELECT id, name, description, budget, status FROM projects WHERE id = $1"

	row := p.db.QueryRowContext(ctx, query, id)

	var project models.Project
	if err := row.Scan(&project.ID, &project.Name, &project.Description, &project.Budget, &project.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return &project, nil
}

// CreateProject inserts a new project into the database
func (p *projectRepository) CreateProject(ctx context.Context, project *models.Project) error {
	query := "INSERT INTO projects (name, description, budget, status) VALUES ($1, $2, $3, $4)"

	_, err := p.db.ExecContext(ctx, query, project.Name, project.Description, project.Budget, project.Status)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProject updates an existing project in the database
func (p *projectRepository) UpdateProject(ctx context.Context, project *models.Project) error {
	query := "UPDATE projects SET name = $1, description = $2, start_date = $3, end_date = $4, budget = $5, status = $6 WHERE id = $7"

	_, err := p.db.ExecContext(ctx, query, project.Name, project.Description, project.Budget, project.Status, project.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProject deletes a project from the database by its ID
func (p *projectRepository) DeleteProject(ctx context.Context, id int64) error {
	query := "DELETE FROM projects WHERE id = $1"

	_, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
