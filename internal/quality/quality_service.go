package quality

import (
	"context"
	"fmt"
	"time"
	models "github.com/BerkatPS/internal"
)

type QualityService interface {
	FindQualityByID(ctx context.Context, id int64) (*models.QualityCheck, error)
	CreateQuality(ctx context.Context, quality *models.QualityCheck) error
	UpdateQuality(ctx context.Context, quality *models.QualityCheck) error
	ShowQualityPerProject(ctx context.Context, projectID int64) ([]models.QualityCheck, error)
	FindQualityByTaskID(ctx context.Context, taskID int64) ([]models.QualityCheck, error)
	FindQualityByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.QualityCheck, error)
	FindQualityIssues(ctx context.Context) ([]models.QualityCheck, error)
	UpdateQualityStatus(ctx context.Context, id int64, status string) error
	FindQualityChecksByInspector(ctx context.Context, inspectorID int64) ([]models.QualityCheck, error)
	FindNonCompliantQualityChecks(ctx context.Context) ([]models.QualityCheck, error)
}

type qualityService struct {
	QualityRepo QualityRepository
}

func NewQualityService(qualityRepo QualityRepository) QualityService {
	return &qualityService{qualityRepo}
}

func (q *qualityService) FindQualityByTaskID(ctx context.Context, taskID int64) ([]models.QualityCheck, error) {
	if taskID <= 0 {
		return nil, fmt.Errorf("invalid task ID")
	}
	qualities, err := q.QualityRepo.FindQualityByTaskID(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve quality: %v", err)
	}
	return qualities, nil
}

func (q *qualityService) FindQualityByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.QualityCheck, error) {
	if startDate.IsZero() || endDate.IsZero() {
		return nil, fmt.Errorf("invalid date range")
	}

	qualities, err := q.QualityRepo.FindQualityByDateRange(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve quality: %v", err)
	}

	return qualities, nil
}

func (q *qualityService) FindQualityIssues(ctx context.Context) ([]models.QualityCheck, error) {
	qualities, err := q.QualityRepo.FindQualityIssues(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve quality: %v", err)
	}
	return qualities, nil
}

func (q *qualityService) UpdateQualityStatus(ctx context.Context, id int64, status string) error {
	if id <= 0 {
		return fmt.Errorf("invalid quality ID")
	}

	if status == "" {
		return fmt.Errorf("status cannot be empty")
	}

	err := q.QualityRepo.UpdateQualityStatus(ctx, id, status)
	if err != nil {
		return fmt.Errorf("failed to update quality: %v", err)
	}

	return nil
}

func (q *qualityService) FindQualityChecksByInspector(ctx context.Context, inspectorID int64) ([]models.QualityCheck, error) {
	if inspectorID <= 0 {
		return nil, fmt.Errorf("invalid inspector ID")
	}

	qualities, err := q.QualityRepo.FindQualityChecksByInspector(ctx, inspectorID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve quality: %v", err)
	}
	return qualities, nil

}

func (q *qualityService) FindNonCompliantQualityChecks(ctx context.Context) ([]models.QualityCheck, error) {
	qualities, err := q.QualityRepo.FindNonCompliantQualityChecks(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve quality: %v", err)
	}
	return qualities, nil
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

func (q *qualityService) FindQualityByID(ctx context.Context, id int64) (*models.QualityCheck, error)  {
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

