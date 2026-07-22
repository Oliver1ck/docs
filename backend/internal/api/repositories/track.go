package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Oliver1ck/docs/internal/api/models"
)

type TrackRepository interface {
	GetByID(ctx context.Context, id int) (*models.Track, error)
	GetByUser(ctx context.Context, userID int, from, to time.Time) ([]models.Track, error)
	GetByStatus(ctx context.Context, status models.TrackStatus) ([]models.Track, error)
	Create(ctx context.Context, track models.Track) (*models.Track, error)
	UpdateStatus(ctx context.Context, id int, status models.TrackStatus) error
	Delete(ctx context.Context, id int) error
	// GetGroups возвращает id групп, привязанных к треку.
	GetGroups(ctx context.Context, trackID int) ([]int, error)
	SetGroups(ctx context.Context, trackID int, groupIDs []int) error
}

type trackRepository struct {
	db *pgxpool.Pool
}

func NewTrackRepository(db *pgxpool.Pool) TrackRepository {
	return &trackRepository{db: db}
}

func (r *trackRepository) GetByID(ctx context.Context, id int) (*models.Track, error) {
	const query = `
		SELECT id, user_id, schedule_rule_id, work_type_id, date, academic_hours, status, comment
		FROM tracks
		WHERE id = $1`

	track, err := scanTrack(r.db.QueryRow(ctx, query, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("trackRepository.GetByID: %w", err)
	}
	return track, nil
}

func (r *trackRepository) GetByUser(ctx context.Context, userID int, from, to time.Time) ([]models.Track, error) {
	const query = `
		SELECT id, user_id, schedule_rule_id, work_type_id, date, academic_hours, status, comment
		FROM tracks
		WHERE user_id = $1 AND date >= $2 AND date <= $3
		ORDER BY date DESC`

	rows, err := r.db.Query(ctx, query, userID, from, to)
	if err != nil {
		return nil, fmt.Errorf("trackRepository.GetByUser: %w", err)
	}
	defer rows.Close()

	return collectTracks(rows)
}

func (r *trackRepository) GetByStatus(ctx context.Context, status models.TrackStatus) ([]models.Track, error) {
	const query = `
		SELECT id, user_id, schedule_rule_id, work_type_id, date, academic_hours, status, comment
		FROM tracks
		WHERE status = $1
		ORDER BY date DESC`

	rows, err := r.db.Query(ctx, query, status)
	if err != nil {
		return nil, fmt.Errorf("trackRepository.GetByStatus: %w", err)
	}
	defer rows.Close()

	return collectTracks(rows)
}

func (r *trackRepository) Create(ctx context.Context, track models.Track) (*models.Track, error) {
	const query = `
		INSERT INTO tracks (user_id, schedule_rule_id, work_type_id, date, academic_hours, status, comment)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, user_id, schedule_rule_id, work_type_id, date, academic_hours, status, comment`

	created, err := scanTrack(r.db.QueryRow(ctx, query,
		track.UserID, track.ScheduleRuleID, track.WorkTypeID,
		track.Date, track.AcademicHours, track.Status, track.Comment,
	))
	if err != nil {
		return nil, fmt.Errorf("trackRepository.Create: %w", err)
	}
	return created, nil
}

func (r *trackRepository) UpdateStatus(ctx context.Context, id int, status models.TrackStatus) error {
	const query = `UPDATE tracks SET status = $2 WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id, status)
	if err != nil {
		return fmt.Errorf("trackRepository.UpdateStatus: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *trackRepository) Delete(ctx context.Context, id int) error {
	const query = `DELETE FROM tracks WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("trackRepository.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *trackRepository) GetGroups(ctx context.Context, trackID int) ([]int, error) {
	const query = `SELECT group_id FROM track_groups WHERE track_id = $1`

	rows, err := r.db.Query(ctx, query, trackID)
	if err != nil {
		return nil, fmt.Errorf("trackRepository.GetGroups: %w", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("trackRepository.GetGroups scan: %w", err)
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}

func (r *trackRepository) SetGroups(ctx context.Context, trackID int, groupIDs []int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("trackRepository.SetGroups begin tx: %w", err)
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx,
		`DELETE FROM track_groups WHERE track_id = $1`, trackID,
	); err != nil {
		return fmt.Errorf("trackRepository.SetGroups delete: %w", err)
	}

	for _, gid := range groupIDs {
		if _, err := tx.Exec(ctx,
			`INSERT INTO track_groups (track_id, group_id) VALUES ($1, $2)`,
			trackID, gid,
		); err != nil {
			return fmt.Errorf("trackRepository.SetGroups insert group %d: %w", gid, err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("trackRepository.SetGroups commit: %w", err)
	}
	return nil
}

func collectTracks(rows pgx.Rows) ([]models.Track, error) {
	var tracks []models.Track
	for rows.Next() {
		t, err := scanTrack(rows)
		if err != nil {
			return nil, fmt.Errorf("track scan: %w", err)
		}
		tracks = append(tracks, *t)
	}
	return tracks, rows.Err()
}

func scanTrack(row pgx.Row) (*models.Track, error) {
	var t models.Track
	err := row.Scan(
		&t.ID, &t.UserID, &t.ScheduleRuleID, &t.WorkTypeID,
		&t.Date, &t.AcademicHours, &t.Status, &t.Comment,
	)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
