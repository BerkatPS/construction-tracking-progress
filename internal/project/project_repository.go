package project

import (
	"database/sql"
	"errors"
	models "github.com/BerkatPS/internal"
)

type ProjectRepository interface {
	FindAll() ([]models.Project, error)
	FindProjectByID(id int64) (*models.Project, error)
	CreateProject(project *models.Project) error
	UpdateProject(project *models.Project) error
	DeleteProject(id int64) error
}

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db}
}

func (p *projectRepository) FindAll() ([]models.Project, error) {
	query := "SELECT id, name, description, start_date, end_date, budget, status FROM projects"

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, &project.StartDate, &project.EndDate, &project.Budget, &project.Status); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (p *projectRepository) FindProjectByID(id int64) (*models.Project, error) {
	query := "SELECT id, name, description, start_date, end_date, budget, status FROM projects WHERE id = $1"

	row := p.db.QueryRow(query, id)

	var project models.Project
	if err := row.Scan(&project.ID, &project.Name, &project.Description, &project.StartDate, &project.EndDate, &project.Budget, &project.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return &project, nil
}

func (p *projectRepository) CreateProject(project *models.Project) error {
	query := "INSERT INTO projects (name, description, start_date, end_date, budget, status) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := p.db.Exec(query, project.Name, project.Description, project.StartDate, project.EndDate, project.Budget, project.Status)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) UpdateProject(project *models.Project) error {
	query := "UPDATE projects SET name = $1, description = $2, start_date = $3, end_date = $4, budget = $5, status = $6 WHERE id = $7"

	_, err := p.db.Exec(query, project.Name, project.Description, project.StartDate, project.EndDate, project.Budget, project.Status, project.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) DeleteProject(id int64) error {
	query := "DELETE FROM projects WHERE id = $1"

	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
