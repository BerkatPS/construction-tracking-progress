package quality

import (
	"context"
	"database/sql"
	"errors"
	"time"

	models "github.com/BerkatPS/internal"
)

type QualityRepository interface {
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

type qualityRepository struct {
	db *sql.DB
}

func NewQualityRepository(db *sql.DB) QualityRepository {
	return &qualityRepository{db}
}

func (q *qualityRepository) FindQualityChecksByInspector(ctx context.Context, inspectorID int64) ([]models.QualityCheck, error) {
	query := "SELECT * FROM quality_checks WHERE inspector_id = $1"

	rows, err := q.db.QueryContext(ctx, query, inspectorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualitys []models.QualityCheck
	for rows.Next() {
		var quality models.QualityCheck
		if err := rows.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
			return nil, err
		}
		qualitys = append(qualitys, quality)
	}

	return qualitys, nil
}

func (q *qualityRepository) FindNonCompliantQualityChecks(ctx context.Context) ([]models.QualityCheck, error) {
	query := "SELECT * FROM quality_checks WHERE status = 'NON_COMPLIANT'"

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualitys []models.QualityCheck
	for rows.Next() {
		var quality models.QualityCheck
		if err := rows.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
			return nil, err
		}
		qualitys = append(qualitys, quality)
	}

	return qualitys, nil
}

func (q *qualityRepository) FindQualityByTaskID(ctx context.Context, taskID int64) ([]models.QualityCheck, error) {
	query := "SELECT * FROM quality_checks WHERE task_id = $1"

	rows, err := q.db.QueryContext(ctx, query, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualitys []models.QualityCheck
	for rows.Next() {
		var quality models.QualityCheck
		if err := rows.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
			return nil, err
		}
		qualitys = append(qualitys, quality)
	}

	return qualitys, nil
}

func (q *qualityRepository) FindQualityByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.QualityCheck, error) {
	query := "SELECT * FROM quality_checks WHERE date BETWEEN $1 AND $2"

	rows, err := q.db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualitys []models.QualityCheck
	for rows.Next() {
		var quality models.QualityCheck
		if err := rows.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
			return nil, err
		}
		qualitys = append(qualitys, quality)
	}

	return qualitys, nil
}

func (q *qualityRepository) FindQualityIssues(ctx context.Context) ([]models.QualityCheck, error) {
	query := "SELECT * FROM quality_checks WHERE status = 'NON_COMPLIANT'"

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualitys []models.QualityCheck
	for rows.Next() {
		var quality models.QualityCheck
		if err := rows.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
			return nil, err
		}
		qualitys = append(qualitys, quality)
	}

	return qualitys, nil
}

func (q *qualityRepository) UpdateQualityStatus(ctx context.Context, id int64, status string) error {
	query := "UPDATE quality_checks SET status = $1 WHERE id = $2"

	_, err := q.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (q *qualityRepository) FindQualityByProjectID(ctx context.Context, projectID int64) ([]models.QualityCheck, error) {
	query := "SELECT * FROM quality_checks WHERE project_id = $1"

	rows, err := q.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualitys []models.QualityCheck
	for rows.Next() {
		var quality models.QualityCheck
		if err := rows.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
			return nil, err
		}
		qualitys = append(qualitys, quality)
	}

	return qualitys, nil
}

func (q *qualityRepository) FindQualityByID(ctx context.Context, id int64) (*models.QualityCheck, error) {
	query := "SELECT * FROM quality_checks WHERE id = $1"

	row := q.db.QueryRowContext(ctx, query, id)
	var quality models.QualityCheck
	if err := row.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("quality not found")
		}
		return nil, err
	}
	return &quality, nil
}

func (q *qualityRepository) CreateQuality(ctx context.Context, quality *models.QualityCheck) error {

	query := "INSERT INTO quality_checks (project_id, inspector_id, date, comments, status) VALUES ($1, $2, $3, $4, $5)"

	_, err := q.db.ExecContext(ctx, query, quality.ProjectID, quality.InspectorID, quality.Date, quality.Comments, quality.Status)
	if err != nil {
		return err
	}
	return nil
}

func (q *qualityRepository) UpdateQuality(ctx context.Context, quality *models.QualityCheck) error {
	query := "UPDATE quality_checks SET project_id = $1, inspector_id = $2, date = $3, comments = $4, status = $5 WHERE id = $6"

	_, err := q.db.ExecContext(ctx, query, quality.ProjectID, quality.InspectorID, quality.Date, quality.Comments, quality.Status, quality.ID)
	if err != nil {
		return err
	}
	return nil
}

func (q *qualityRepository) ShowQualityPerProject(ctx context.Context, projectID int64) ([]models.QualityCheck, error) {

	query := "SELECT * FROM quality_checks WHERE project_id = $1 JOIN users ON quality_checks.inspector_id = users.id ORDER BY date DESC"

	rows, err := q.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualitys []models.QualityCheck
	for rows.Next() {
		var quality models.QualityCheck
		if err := rows.Scan(&quality.ID, &quality.ProjectID, &quality.InspectorID, &quality.Date, &quality.Comments, &quality.Status); err != nil {
			return nil, err
		}
		qualitys = append(qualitys, quality)
	}

	return qualitys, nil
}
