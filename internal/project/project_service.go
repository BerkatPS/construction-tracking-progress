package project

import (
	"context"
	"errors"
	models "github.com/BerkatPS/internal"
)

type ProjectService interface {
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

type projectService struct {
	ProjectRepo ProjectRepository
}

// NewProjectService creates a new instance of ProjectService
func NewProjectService(ProjectRepo ProjectRepository) ProjectService {
	return &projectService{ProjectRepo}
}

func (p *projectService) UpdateProjectBudget(ctx context.Context, projectId int64, newBudget float64) error {
	if projectId <= 0 {
		return errors.New("invalid project ID")
	}

	if newBudget <= 0 {
		return errors.New("invalid new budget")
	}

	if err := p.ProjectRepo.UpdateProjectBudget(ctx, projectId, newBudget); err != nil {
		return err
	}

	return nil
}

func (p *projectService) DeleteProjectDocument(ctx context.Context, documentId int64) error {
	if documentId <= 0 {
		return errors.New("invalid document ID")
	}

	if err := p.ProjectRepo.DeleteProjectDocument(ctx, documentId); err != nil {
		return err
	}

	return nil
}

func (p *projectService) UploadProjectDocument(ctx context.Context, projectId int64, document *models.Document) error {
	if projectId <= 0 {
		return errors.New("invalid project ID")
	}

	if document.Name == "" || document.Type == "" || document.URL == "" {
		return errors.New("missing required document fields")
	}

	if err := p.ProjectRepo.UploadProjectDocument(ctx, projectId, document); err != nil {
		return err
	}

	return nil
}

func (p *projectService) TrackProjectExpenses(ctx context.Context, expense *models.Expense) error {
	if expense.ProjectID <= 0 {
		return errors.New("invalid project ID")
	}

	if expense.Amount <= 0 {
		return errors.New("invalid expense amount")
	}

	if expense.Description == "" {
		return errors.New("missing required expense fields")
	}

	if err := p.ProjectRepo.TrackProjectExpenses(ctx, expense); err != nil {
		return err
	}

	return nil
}

func (p *projectService) FindExpensesByProject(ctx context.Context, projectId int64) ([]models.Expense, error) {
	if projectId <= 0 {
		return nil, errors.New("invalid project ID")
	}

	expenses, err := p.ProjectRepo.FindExpensesByProject(ctx, projectId)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (p *projectService) UpdateProjectStatus(ctx context.Context, id int64, status string) error {
	if id <= 0 {
		return errors.New("invalid project ID")
	}

	if status == "" {
		return errors.New("missing required project fields")
	}

	if err := p.ProjectRepo.UpdateProjectStatus(ctx, id, status); err != nil {
		return err
	}

	return nil
}

func (p *projectService) AddTeamMemberToProject(ctx context.Context, projectId int64, userId int64) error {
	if projectId <= 0 {
		return errors.New("invalid project ID")
	}

	if userId <= 0 {
		return errors.New("invalid user ID")
	}

	if err := p.ProjectRepo.AddTeamMemberToProject(ctx, projectId, userId); err != nil {
		return err
	}

	return nil
}

func (p *projectService) RemoveTeamMemberFromProject(ctx context.Context, projectId int64, userId int64) error {
	if projectId <= 0 {
		return errors.New("invalid project ID")
	}

	if userId <= 0 {
		return errors.New("invalid user ID")
	}

	if err := p.ProjectRepo.RemoveTeamMemberFromProject(ctx, projectId, userId); err != nil {
		return err
	}

	return nil
}

func (p *projectService) UpdateProjectTeamRole(ctx context.Context, projectId int64, userId int64, role string) error {
	if projectId <= 0 {
		return errors.New("invalid project ID")
	}

	if userId <= 0 {
		return errors.New("invalid user ID")
	}

	if role == "" {
		return errors.New("missing required project fields")
	}

	if err := p.ProjectRepo.UpdateProjectTeamRole(ctx, projectId, userId, role); err != nil {
		return err
	}

	return nil
}

func (p *projectService) FindProjectsByStatus(ctx context.Context, status string) ([]models.Project, error) {
	if status == "" {
		return nil, errors.New("missing required project fields")
	}

	projects, err := p.ProjectRepo.FindProjectsByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	return projects, nil
}



// FindAll retrieves all projects from the repository
func (p *projectService) FindAll(ctx context.Context) ([]models.Project, error) {
	projects, err := p.ProjectRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// FindProjectByID retrieves a project by its ID
func (p *projectService) FindProjectByID(ctx context.Context, id int64) (*models.Project, error) {
	if id <= 0 {
		return nil, errors.New("invalid project ID")
	}

	project, err := p.ProjectRepo.FindProjectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// CreateProject validates and creates a new project
func (p *projectService) CreateProject(ctx context.Context, project *models.Project) error {
	if project.Name == "" {
		return errors.New("missing required project fields")
	}

	if err := p.ProjectRepo.CreateProject(ctx, project); err != nil {
		return err
	}

	return nil
}

// UpdateProject validates and updates an existing project
func (p *projectService) UpdateProject(ctx context.Context, project *models.Project) error {
	if project.ID <= 0 {
		return errors.New("invalid project ID")
	}

	if project.Name == "" {
		return errors.New("missing required project fields")
	}

	existingProject, err := p.ProjectRepo.FindProjectByID(ctx, project.ID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("project not found")
	}

	if err := p.ProjectRepo.UpdateProject(ctx, project); err != nil {
		return err
	}

	return nil
}

// DeleteProject deletes a project by its ID
func (p *projectService) DeleteProject(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid project ID")
	}

	existingProject, err := p.ProjectRepo.FindProjectByID(ctx, id)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("project not found")
	}

	if err := p.ProjectRepo.DeleteProject(ctx, id); err != nil {
		return err
	}

	return nil
}
