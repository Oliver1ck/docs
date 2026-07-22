package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Oliver1ck/docs/internal/api/models"
)

var ErrNotFound = errors.New("not found")

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user models.User) (*models.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	const query = `
		SELECT id, username, email, password, role_id, created_at, updated_at
		FROM users
		WHERE id = $1`

	user, err := scanUser(r.db.QueryRow(ctx, query, id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByID: %w", err)
	}
	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	const query = `
		SELECT id, username, email, password, role_id, created_at, updated_at
		FROM users
		WHERE email = $1`

	user, err := scanUser(r.db.QueryRow(ctx, query, email))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetByEmail: %w", err)
	}
	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user models.User) (*models.User, error) {
	const query = `
		INSERT INTO users (username, email, password, role_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, username, email, password, role_id, created_at, updated_at`

	created, err := scanUser(r.db.QueryRow(ctx, query,
		user.Username,
		user.Email,
		user.Password,
		user.RoleID,
	))
	if err != nil {
		return nil, fmt.Errorf("userRepository.Create: %w", err)
	}
	return created, nil
}

func scanUser(row pgx.Row) (*models.User, error) {
	var u models.User
	err := row.Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.Password,
		&u.RoleID,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
