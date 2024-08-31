package presence

import (
	"context"
	"database/sql"
	models "github.com/BerkatPS/internal"
)

type PresenceRepository interface {
	FindAll(ctx context.Context) ([]models.Presence, error)
	FindPresenceByID(ctx context.Context, id int64) (*models.Presence, error)
	FindPresenceByUserID(ctx context.Context, userID int64) (*models.Presence, error)
	CreatePresence(ctx context.Context, presence *models.Presence) error
	UpdatePresence(ctx context.Context, presence *models.Presence) error
}

type presenceRepository struct {
	db *sql.DB
}

func NewPresenceRepository(db *sql.DB) PresenceRepository {
	return &presenceRepository{db}
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
	query := "SELECT id, user_id, status FROM presences WHERE user_id = $1"

	var presence models.Presence
	err := p.db.QueryRowContext(ctx, query, userID).Scan(&presence.ID, &presence.UserID, &presence.Status)
	if err != nil {
		return nil, err
	}

	return &presence, nil
}

func (p *presenceRepository) CreatePresence(ctx context.Context, presence *models.Presence) error {
	query := "INSERT INTO presences (user_id, status) VALUES ($1, $2)"

	_, err := p.db.ExecContext(ctx, query, presence.UserID, presence.Status)
	return err
}

func (p *presenceRepository) UpdatePresence(ctx context.Context, presence *models.Presence) error {
	query := "UPDATE presences SET status = $1 WHERE id = $2"

	_, err := p.db.ExecContext(ctx, query, presence.Status, presence.ID)
	return err
}
