package quality

import (
	"context"
	"database/sql"
	"errors"
	models "github.com/BerkatPS/internal"
)

type QualityRepository interface {
	FindQualityByID(ctx context.Context, id int64) (*models.QualityCheck, error)
	CreateQuality(ctx context.Context, quality *models.QualityCheck) error
	UpdateQuality(ctx context.Context, quality *models.QualityCheck) error
	ShowQualityPerProject(ctx context.Context, projectID int64) ([]models.QualityCheck, error)
}

type qualityRepository struct {
	db *sql.DB
}

func NewQualityRepository(db *sql.DB) QualityRepository {
	return &qualityRepository{db}
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
