package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Oliver1ck/docs/internal/api/models"
	"github.com/Oliver1ck/docs/internal/api/repositories"
)

var (
	ErrTrackNotFound   = errors.New("track not found")
	ErrTrackForbidden  = errors.New("track belongs to another user")
	ErrInvalidStatus   = errors.New("invalid status transition")
)

type CreateTrackInput struct {
	ScheduleRuleID *int
	WorkTypeID     int
	Date           time.Time
	AcademicHours  float64
	Comment        string
	GroupIDs       []int
}

type TrackService interface {
	GetByID(ctx context.Context, id int) (*models.Track, error)
	GetByUser(ctx context.Context, userID int, from, to time.Time) ([]models.Track, error)
	GetByStatus(ctx context.Context, status models.TrackStatus) ([]models.Track, error)
	Create(ctx context.Context, userID int, input CreateTrackInput) (*models.Track, error)
	Confirm(ctx context.Context, id int) error
	Reject(ctx context.Context, id int) error
	Delete(ctx context.Context, userID, id int) error
}

type trackService struct {
	repo repositories.TrackRepository
}

func NewTrackService(repo repositories.TrackRepository) TrackService {
	return &trackService{repo: repo}
}

func (s *trackService) GetByID(ctx context.Context, id int) (*models.Track, error) {
	track, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, repositories.ErrNotFound) {
		return nil, ErrTrackNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("trackService.GetByID: %w", err)
	}
	return track, nil
}

func (s *trackService) GetByUser(ctx context.Context, userID int, from, to time.Time) ([]models.Track, error) {
	tracks, err := s.repo.GetByUser(ctx, userID, from, to)
	if err != nil {
		return nil, fmt.Errorf("trackService.GetByUser: %w", err)
	}
	return tracks, nil
}

func (s *trackService) GetByStatus(ctx context.Context, status models.TrackStatus) ([]models.Track, error) {
	tracks, err := s.repo.GetByStatus(ctx, status)
	if err != nil {
		return nil, fmt.Errorf("trackService.GetByStatus: %w", err)
	}
	return tracks, nil
}

func (s *trackService) Create(ctx context.Context, userID int, input CreateTrackInput) (*models.Track, error) {
	track, err := s.repo.Create(ctx, models.Track{
		UserID:         userID,
		ScheduleRuleID: input.ScheduleRuleID,
		WorkTypeID:     input.WorkTypeID,
		Date:           input.Date,
		AcademicHours:  input.AcademicHours,
		Status:         models.TrackStatusPending,
		Comment:        input.Comment,
	})
	if err != nil {
		return nil, fmt.Errorf("trackService.Create: %w", err)
	}

	if len(input.GroupIDs) > 0 {
		if err := s.repo.SetGroups(ctx, track.ID, input.GroupIDs); err != nil {
			return nil, fmt.Errorf("trackService.Create set groups: %w", err)
		}
	}

	return track, nil
}

func (s *trackService) Confirm(ctx context.Context, id int) error {
	track, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, repositories.ErrNotFound) {
		return ErrTrackNotFound
	}
	if err != nil {
		return fmt.Errorf("trackService.Confirm get: %w", err)
	}
	if track.Status != models.TrackStatusPending {
		return ErrInvalidStatus
	}
	if err := s.repo.UpdateStatus(ctx, id, models.TrackStatusConfirmed); err != nil {
		return fmt.Errorf("trackService.Confirm update: %w", err)
	}
	return nil
}

func (s *trackService) Reject(ctx context.Context, id int) error {
	track, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, repositories.ErrNotFound) {
		return ErrTrackNotFound
	}
	if err != nil {
		return fmt.Errorf("trackService.Reject get: %w", err)
	}
	if track.Status != models.TrackStatusPending {
		return ErrInvalidStatus
	}
	if err := s.repo.UpdateStatus(ctx, id, models.TrackStatusRejected); err != nil {
		return fmt.Errorf("trackService.Reject update: %w", err)
	}
	return nil
}

func (s *trackService) Delete(ctx context.Context, userID, id int) error {
	track, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, repositories.ErrNotFound) {
		return ErrTrackNotFound
	}
	if err != nil {
		return fmt.Errorf("trackService.Delete get: %w", err)
	}
	if track.UserID != userID {
		return ErrTrackForbidden
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("trackService.Delete: %w", err)
	}
	return nil
}
