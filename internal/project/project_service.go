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
}

type projectService struct {
	ProjectRepo ProjectRepository
}

// NewProjectService creates a new instance of ProjectService
func NewProjectService(ProjectRepo ProjectRepository) ProjectService {
	return &projectService{ProjectRepo}
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
