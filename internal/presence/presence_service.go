package presence

import (
	"context"
	"errors"
	models "github.com/BerkatPS/internal"
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
