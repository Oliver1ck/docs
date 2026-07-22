package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Oliver1ck/docs/internal/config"
)

func New(ctx context.Context, cfg config.Postgres) (*pgxpool.Pool, error) {
	connStr := cfg.CreateDBConnectionString()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to create pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("postgres: failed to ping: %w", err)
	}

	return pool, nil
}
