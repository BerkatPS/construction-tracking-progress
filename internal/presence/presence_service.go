package presence

import (
	"context"
	"errors"
	"fmt"
	models "github.com/BerkatPS/internal"
	"time"
)

type PresenceService interface {
	FindAll(ctx context.Context) ([]models.Presence, error)
	FindPresenceByID(ctx context.Context, id int64) (*models.Presence, error)
	FindPresenceByUserID(ctx context.Context, userID int64) (*models.Presence, error)
	CreatePresence(ctx context.Context, presence *models.Presence) error
	UpdatePresence(ctx context.Context, presence *models.Presence) error
}

type presenceService struct {
	presenceRepository PresenceRepository
}

func NewPresenceService(presenceRepository PresenceRepository) PresenceService {
	return &presenceService{presenceRepository}
}

func (p *presenceService) FindAll(ctx context.Context) ([]models.Presence, error) {
	return p.presenceRepository.FindAll(ctx)
}

func (p *presenceService) FindPresenceByID(ctx context.Context, id int64) (*models.Presence, error) {
	if id <= 0 {
		return nil, errors.New("invalid presence ID")
	}
	return p.presenceRepository.FindPresenceByID(ctx, id)
}

func (p *presenceService) FindPresenceByUserID(ctx context.Context, userID int64) (*models.Presence, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}

	return p.presenceRepository.FindPresenceByUserID(ctx, userID)
}

func (p *presenceService) CreatePresence(ctx context.Context, presence *models.Presence) error {
	if presence.UserID <= 0 {
		return errors.New("invalid user ID")
	}
	if presence.Comments == "" {
		return errors.New("comments are required")
	}
	if presence.Status == "" {
		return errors.New("status is required")
	}

	today := time.Now().Format("2006-01-02")

	existingPresence, err := p.presenceRepository.FindPresenceByUserIDAndDate(ctx, presence.UserID, today)
	if err != nil {
		return fmt.Errorf("failed to check if presence already exists: %w", err)
	}

	if existingPresence != nil {
		return errors.New("presence already exists")
	}
	presence.Date = time.Now()
	return p.presenceRepository.CreatePresence(ctx, presence)
}

func (p *presenceService) UpdatePresence(ctx context.Context, presence *models.Presence) error {
	if presence.UserID <= 0 {
		return errors.New("invalid user ID")
	}

	if presence.Status == "" {
		return errors.New("status is required")
	}

	return p.presenceRepository.UpdatePresence(ctx, presence)
}
