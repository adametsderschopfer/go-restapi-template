package postgresql

import "C"
import (
	"app/internal/config"
	"app/internal/repository"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(connectionString string) (*Repository, error) {
	const op = "repository.postgresql.New"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Repository{db: db}, nil
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

/*
	Example methods
*/

func (s *Repository) SaveUrl(urlToSave string, alias string) (int64, error) {
	const op = "repository.postgresql.SaveUrl"

	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES(?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Repository) GetUrlByAlias(alias string) (string, error) {
	const op = "repository.postgresql.GetUrl"

	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var resUrl string
	err = stmt.QueryRow(alias).Scan(&resUrl)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repository.ErrURLNotFound
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resUrl, nil
}
