package quality

import (
	"context"
	"fmt"
	models "github.com/BerkatPS/internal"
)

type QualityService interface {
	FindQualityByID(ctx context.Context, id int64) (*models.QualityCheck, error)
	CreateQuality(ctx context.Context, quality *models.QualityCheck) error
	UpdateQuality(ctx context.Context, quality *models.QualityCheck) error
	ShowQualityPerProject(ctx context.Context, projectID int64) ([]models.QualityCheck, error)
}

type qualityService struct {
	QualityRepo QualityRepository
}

func NewQualityService(qualityRepo QualityRepository) QualityService {
	return &qualityService{qualityRepo}
}

func (q *qualityService) FindQualityByID(ctx context.Context, id int64) (*models.QualityCheck, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid quality ID")
	}
	quality, err := q.QualityRepo.FindQualityByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve quality: %v", err)
	}
	return quality, nil

}

func (q *qualityService) CreateQuality(ctx context.Context, quality *models.QualityCheck) error {
	if quality.Comments == "" {
		return fmt.Errorf("comments cannot be empty")
	}

	if quality.Date.IsZero() {
		return fmt.Errorf("date cannot be empty")
	}

	if quality.InspectorID <= 0 {
		return fmt.Errorf("invalid inspector ID")
	}

	if quality.ProjectID <= 0 {
		return fmt.Errorf("invalid project ID")
	}

	if quality.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}

	err := q.QualityRepo.CreateQuality(ctx, quality)
	if err != nil {
		return fmt.Errorf("failed to create quality: %v", err)
	}
	return nil
}

func (q *qualityService) UpdateQuality(ctx context.Context, quality *models.QualityCheck) error {

	if quality.ID <= 0 {
		return fmt.Errorf("invalid quality ID")
	}

	if quality.Comments == "" {
		return fmt.Errorf("comments cannot be empty")
	}

	if quality.Date.IsZero() {
		return fmt.Errorf("date cannot be empty")
	}

	if quality.InspectorID <= 0 {
		return fmt.Errorf("invalid inspector ID")
	}

	if quality.ProjectID <= 0 {
		return fmt.Errorf("invalid project ID")
	}

	if quality.Status == "" {
		return fmt.Errorf("status cannot be empty")
	}

	err := q.QualityRepo.UpdateQuality(ctx, quality)
	if err != nil {
		return fmt.Errorf("failed to update quality: %v", err)
	}
	return nil
}

func (q *qualityService) ShowQualityPerProject(ctx context.Context, projectID int64) ([]models.QualityCheck, error) {

	if projectID <= 0 {
		return nil, fmt.Errorf("invalid project ID")
	}

	qualities, err := q.QualityRepo.ShowQualityPerProject(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve quality: %v", err)
	}
	return qualities, nil

}
