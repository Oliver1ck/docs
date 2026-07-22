package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Oliver1ck/docs/internal/api/models"
)

type GroupRepository interface {
	GetAll(ctx context.Context) ([]models.Group, error)
	GetByID(ctx context.Context, id int) (*models.Group, error)
	Create(ctx context.Context, name string) (*models.Group, error)
	Delete(ctx context.Context, id int) error
}

type groupRepository struct {
	db *pgxpool.Pool
}

func NewGroupRepository(db *pgxpool.Pool) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) GetAll(ctx context.Context) ([]models.Group, error) {
	const query = `
		SELECT id, name, created_at
		FROM groups
		ORDER BY name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("groupRepository.GetAll: %w", err)
	}
	defer rows.Close()

	var groups []models.Group
	for rows.Next() {
		var g models.Group
		if err := rows.Scan(&g.ID, &g.Name, &g.CreatedAt); err != nil {
			return nil, fmt.Errorf("groupRepository.GetAll scan: %w", err)
		}
		groups = append(groups, g)
	}
	return groups, rows.Err()
}

func (r *groupRepository) GetByID(ctx context.Context, id int) (*models.Group, error) {
	const query = `
		SELECT id, name, created_at
		FROM groups
		WHERE id = $1`

	var g models.Group
	err := r.db.QueryRow(ctx, query, id).Scan(&g.ID, &g.Name, &g.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("groupRepository.GetByID: %w", err)
	}
	return &g, nil
}

func (r *groupRepository) Create(ctx context.Context, name string) (*models.Group, error) {
	const query = `
		INSERT INTO groups (name)
		VALUES ($1)
		RETURNING id, name, created_at`

	var g models.Group
	err := r.db.QueryRow(ctx, query, name).Scan(&g.ID, &g.Name, &g.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("groupRepository.Create: %w", err)
	}
	return &g, nil
}

func (r *groupRepository) Delete(ctx context.Context, id int) error {
	const query = `DELETE FROM groups WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("groupRepository.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
