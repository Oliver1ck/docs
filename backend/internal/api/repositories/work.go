package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Oliver1ck/docs/internal/api/models"
)

type WorkRepository interface {
	GetAllCategories(ctx context.Context) ([]models.WorkCategory, error)

	GetAllTypes(ctx context.Context) ([]models.WorkType, error)
	GetTypeByID(ctx context.Context, id int) (*models.WorkType, error)
	GetTypesByCategory(ctx context.Context, categoryID int) ([]models.WorkType, error)
	CreateType(ctx context.Context, wt models.WorkType) (*models.WorkType, error)
	SetTypeActive(ctx context.Context, id int, active bool) error
}

type workRepository struct {
	db *pgxpool.Pool
}

func NewWorkRepository(db *pgxpool.Pool) WorkRepository {
	return &workRepository{db: db}
}

// --- WorkCategory ---

func (r *workRepository) GetAllCategories(ctx context.Context) ([]models.WorkCategory, error) {
	const query = `
		SELECT id, name, sort_order
		FROM work_categories
		ORDER BY sort_order, name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("workRepository.GetAllCategories: %w", err)
	}
	defer rows.Close()

	var cats []models.WorkCategory
	for rows.Next() {
		var c models.WorkCategory
		if err := rows.Scan(&c.ID, &c.Name, &c.SortOrder); err != nil {
			return nil, fmt.Errorf("workRepository.GetAllCategories scan: %w", err)
		}
		cats = append(cats, c)
	}
	return cats, rows.Err()
}

// --- WorkType ---

func (r *workRepository) GetAllTypes(ctx context.Context) ([]models.WorkType, error) {
	const query = `
		SELECT id, name, short_name, category_id, is_active
		FROM work_types
		ORDER BY category_id, name`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("workRepository.GetAllTypes: %w", err)
	}
	defer rows.Close()

	var types []models.WorkType
	for rows.Next() {
		var wt models.WorkType
		if err := rows.Scan(&wt.ID, &wt.Name, &wt.ShortName, &wt.CategoryID, &wt.IsActive); err != nil {
			return nil, fmt.Errorf("workRepository.GetAllTypes scan: %w", err)
		}
		types = append(types, wt)
	}
	return types, rows.Err()
}

func (r *workRepository) GetTypeByID(ctx context.Context, id int) (*models.WorkType, error) {
	const query = `
		SELECT id, name, short_name, category_id, is_active
		FROM work_types
		WHERE id = $1`

	var wt models.WorkType
	err := r.db.QueryRow(ctx, query, id).Scan(&wt.ID, &wt.Name, &wt.ShortName, &wt.CategoryID, &wt.IsActive)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("workRepository.GetTypeByID: %w", err)
	}
	return &wt, nil
}

func (r *workRepository) GetTypesByCategory(ctx context.Context, categoryID int) ([]models.WorkType, error) {
	const query = `
		SELECT id, name, short_name, category_id, is_active
		FROM work_types
		WHERE category_id = $1
		ORDER BY name`

	rows, err := r.db.Query(ctx, query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("workRepository.GetTypesByCategory: %w", err)
	}
	defer rows.Close()

	var types []models.WorkType
	for rows.Next() {
		var wt models.WorkType
		if err := rows.Scan(&wt.ID, &wt.Name, &wt.ShortName, &wt.CategoryID, &wt.IsActive); err != nil {
			return nil, fmt.Errorf("workRepository.GetTypesByCategory scan: %w", err)
		}
		types = append(types, wt)
	}
	return types, rows.Err()
}

func (r *workRepository) CreateType(ctx context.Context, wt models.WorkType) (*models.WorkType, error) {
	const query = `
		INSERT INTO work_types (name, short_name, category_id, is_active)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, short_name, category_id, is_active`

	var created models.WorkType
	err := r.db.QueryRow(ctx, query, wt.Name, wt.ShortName, wt.CategoryID, wt.IsActive).
		Scan(&created.ID, &created.Name, &created.ShortName, &created.CategoryID, &created.IsActive)
	if err != nil {
		return nil, fmt.Errorf("workRepository.CreateType: %w", err)
	}
	return &created, nil
}

func (r *workRepository) SetTypeActive(ctx context.Context, id int, active bool) error {
	const query = `UPDATE work_types SET is_active = $2 WHERE id = $1`

	tag, err := r.db.Exec(ctx, query, id, active)
	if err != nil {
		return fmt.Errorf("workRepository.SetTypeActive: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
