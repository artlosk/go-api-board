package database

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func OpenPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping: %w", err)
	}

	if err := runMigrations(ctx, pool); err != nil {
		pool.Close()
		return nil, fmt.Errorf("migrations: %w", err)
	}

	return pool, nil
}

func runMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	data, err := migrationsFS.ReadFile("migrations/001_init.sql")
	if err != nil {
		return err
	}
	_, err = pool.Exec(ctx, string(data))
	if err != nil {
		return err
	}
	log.Println("migrations applied: 001_init.sql")
	return nil
}
