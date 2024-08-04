package postgresql

import (
	"app/internal/repository"
	"database/sql"
)

func NewPostgresqlRepositories(db *sql.DB) *repository.Repositories {
	return &repository.Repositories{
		Example: NewExampleRepo(db),
	}
}
