package project

import models "github.com/BerkatPS/internal"

type ProjectService interface {
	FindAll() ([]models.Project, error)
	FindProjectByID(id int64) (*models.Project, error)
	CreateProject(project *models.Project) error
	UpdateProject(project *models.Project) error
	DeleteProject(id int64) error
}

type projectService struct {
	ProjectRepo ProjectRepository
}

func NewProjectService(ProjectRepo ProjectRepository) ProjectService {
	return &projectService{ProjectRepo}
}

func (p *projectService) FindAll() ([]models.Project, error) {
	return p.ProjectRepo.FindAll()
}

func (p *projectService) FindProjectByID(id int64) (*models.Project, error) {
	return p.ProjectRepo.FindProjectByID(id)
}

func (p *projectService) CreateProject(project *models.Project) error {
	return p.ProjectRepo.CreateProject(project)
}

func (p *projectService) UpdateProject(project *models.Project) error {
	return p.ProjectRepo.UpdateProject(project)
}

func (p *projectService) DeleteProject(id int64) error {
	return p.ProjectRepo.DeleteProject(id)
}
