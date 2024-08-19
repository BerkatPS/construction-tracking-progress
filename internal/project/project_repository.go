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
	// FindProjectsByStatus allows filtering projects by their status (e.g., ongoing, completed, delayed)
	FindProjectsByStatus(ctx context.Context, status string) ([]models.Project, error)
	// UpdateProjectStatus updates the status of a project, which is crucial for real-time monitoring
	UpdateProjectStatus(ctx context.Context, id int64, status string) error
	// AddTeamMemberToProject adds a new team member to an existing project to improve team collaboration
	AddTeamMemberToProject(ctx context.Context, projectId int64, userId int64) error
	// RemoveTeamMemberFromProject removes a team member from a project, useful for managing team composition
	RemoveTeamMemberFromProject(ctx context.Context, projectId int64, userId int64) error
	// UpdateProjectTeamRole updates the role of a team member in a project
	UpdateProjectTeamRole(ctx context.Context, projectId int64, userId int64, role string) error
	// TrackProjectExpenses tracks and logs an expense related to a project, aiding in real-time budget management
	TrackProjectExpenses(ctx context.Context, expense *models.Expense) error
	// FindExpensesByProject finds all expenses associated with a project
	FindExpensesByProject(ctx context.Context, projectId int64) ([]models.Expense, error)
	// UpdateProjectBudget allows updating the overall budget for a project, useful for real-time adjustments
	UpdateProjectBudget(ctx context.Context, projectId int64, newBudget float64) error
	// DeleteProjectDocument deletes a document from a project's records
	DeleteProjectDocument(ctx context.Context, documentId int64) error
	// UploadProjectDocument uploads a document related to a project, storing it for easy access
	UploadProjectDocument(ctx context.Context, projectId int64, document *models.Document) error
}

type projectRepository struct {
	db *sql.DB
}

// NewProjectRepository creates a new instance of ProjectRepository
func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db}
}

func (p *projectRepository) FindExpensesByProject(ctx context.Context, projectId int64) ([]models.Expense, error) {
	query := "SELECT id, project_id, amount, description, date FROM expenses WHERE project_id = $1"

	rows, err := p.db.QueryContext(ctx, query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.ProjectID, &expense.Amount, &expense.Description, &expense.Date); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return expenses, nil
}

func (p *projectRepository) UpdateProjectBudget(ctx context.Context, projectId int64, newBudget float64) error {
	query := "UPDATE projects SET budget = $1 WHERE id = $2"

	_, err := p.db.ExecContext(ctx, query, newBudget, projectId)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) DeleteProjectDocument(ctx context.Context, documentId int64) error {
	query := "DELETE FROM documents WHERE id = $1"

	_, err := p.db.ExecContext(ctx, query, documentId)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) UploadProjectDocument(ctx context.Context, projectId int64, document *models.Document) error {
	query := "INSERT INTO documents (project_id, name, type, url) VALUES ($1, $2, $3, $4)"

	_, err := p.db.ExecContext(ctx, query, projectId, document.Name, document.Type, document.URL)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) TrackProjectExpenses(ctx context.Context, expense *models.Expense) error {
	query := "INSERT INTO expenses (project_id, amount, description, date) VALUES ($1, $2, $3, $4)"

	_, err := p.db.ExecContext(ctx, query, expense.ProjectID, expense.Amount, expense.Description, expense.Date)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) UpdateProjectStatus(ctx context.Context, id int64, status string) error {
	query := "UPDATE projects SET status = $1 WHERE id = $2"

	_, err := p.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) AddTeamMemberToProject(ctx context.Context, projectId int64, userId int64) error {
	query := "INSERT INTO project_team (project_id, user_id) VALUES ($1, $2)"

	_, err := p.db.ExecContext(ctx, query, projectId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) RemoveTeamMemberFromProject(ctx context.Context, projectId int64, userId int64) error {
	query := "DELETE FROM project_team WHERE project_id = $1 AND user_id = $2"

	_, err := p.db.ExecContext(ctx, query, projectId, userId)

	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) UpdateProjectTeamRole(ctx context.Context, projectId int64, userId int64, role string) error {
	query := "UPDATE project_team SET role = $1 WHERE project_id = $2 AND user_id = $3"

	_, err := p.db.ExecContext(ctx, query, role, projectId, userId)

	if err != nil {
		return err
	}
	return nil
}

func (p *projectRepository) FindProjectsByStatus(ctx context.Context, status string) ([]models.Project, error) {
	query := "SELECT id, name, description, budget, status FROM projects WHERE status = $1"

	rows, err := p.db.QueryContext(ctx, query, status)
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
