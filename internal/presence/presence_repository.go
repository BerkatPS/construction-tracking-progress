package presence

import (
	"context"
	"database/sql"
	"fmt"
	models "github.com/BerkatPS/internal"
)

type PresenceRepository interface {
	FindAll(ctx context.Context) ([]models.Presence, error)
	FindPresenceByID(ctx context.Context, id int64) (*models.Presence, error)
	FindPresenceByUserID(ctx context.Context, userID int64) (*models.Presence, error)
	CreatePresence(ctx context.Context, presence *models.Presence) error
	FindPresenceByUserIDAndDate(ctx context.Context, userID int64, date string) (*models.Presence, error)
	UpdatePresence(ctx context.Context, presence *models.Presence) error
}

type presenceRepository struct {
	db *sql.DB
}

func NewPresenceRepository(db *sql.DB) PresenceRepository {
	return &presenceRepository{db}
}

func (p *presenceRepository) FindPresenceByUserIDAndDate(ctx context.Context, userID int64, date string) (*models.Presence, error) {
	query := "SELECT id, user_id, date FROM presences WHERE user_id = $1 AND DATE(date) = $2"

	var presence models.Presence
	err := p.db.QueryRowContext(ctx, query, userID, date).Scan(&presence.ID, &presence.UserID, &presence.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			// No presence found for the given user and date, return nil without error
			return nil, nil
		}
		return nil, err
	}

	return &presence, nil
}

func (p *presenceRepository) FindAll(ctx context.Context) ([]models.Presence, error) {
	query := "SELECT id, user_id, status FROM presences"

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var presences []models.Presence
	for rows.Next() {
		var presence models.Presence
		if err := rows.Scan(&presence.ID, &presence.UserID, &presence.Status); err != nil {
			return nil, err
		}
		presences = append(presences, presence)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return presences, nil
}

func (p *presenceRepository) FindPresenceByID(ctx context.Context, id int64) (*models.Presence, error) {
	query := "SELECT id, user_id, status FROM presences WHERE id = $1"

	var presence models.Presence
	err := p.db.QueryRowContext(ctx, query, id).Scan(&presence.ID, &presence.UserID, &presence.Status)
	if err != nil {
		return nil, err
	}

	return &presence, nil
}

func (p *presenceRepository) FindPresenceByUserID(ctx context.Context, userID int64) (*models.Presence, error) {
	query := "SELECT id, user_id, status, comments, date FROM presences WHERE user_id = $1"

	var presence models.Presence
	err := p.db.QueryRowContext(ctx, query, userID).Scan(&presence.ID, &presence.UserID, &presence.Status, &presence.Comments, &presence.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("presence not found for user ID %d", userID)
		}

		return nil, err
	}

	return &presence, nil
}

func (p *presenceRepository) CreatePresence(ctx context.Context, presence *models.Presence) error {
	query := "INSERT INTO presences (user_id, status, comments, date) VALUES ($1, $2, $3, $4)"

	_, err := p.db.ExecContext(ctx, query, presence.UserID, presence.Status, presence.Comments, presence.Date)
	return err
}

func (p *presenceRepository) UpdatePresence(ctx context.Context, presence *models.Presence) error {
	query := "UPDATE presences SET status = $1 WHERE id = $2"

	_, err := p.db.ExecContext(ctx, query, presence.Status, presence.ID)
	return err
}
