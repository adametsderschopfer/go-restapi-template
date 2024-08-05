package postgresql

import "C"
import (
	"app/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func New(connectionString string) (*sql.DB, error) {
	const op = "database.postgresql.New"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

func CreateConnectionString(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)
}
