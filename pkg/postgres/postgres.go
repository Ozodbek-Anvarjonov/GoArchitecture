package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Postgres driver
)

type Postgres struct {
	DB *sql.DB
}

// New yaratadi va Postgresga ulanadi
func New(dsn string, maxPoolSize int) (*Postgres, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	db.SetMaxOpenConns(maxPoolSize)
	db.SetMaxIdleConns(maxPoolSize / 2)
	db.SetConnMaxLifetime(time.Hour)

	// Ping orqali aloqani tekshirish (5 soniya ichida)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &Postgres{DB: db}, nil
}

// Close ulanishni to'g'ri yopadi
func (p *Postgres) Close() error {
	return p.DB.Close()
}
