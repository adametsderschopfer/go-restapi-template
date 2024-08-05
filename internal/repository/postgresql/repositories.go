package postgresql

import (
	"app/internal/repository/repos"
	"database/sql"
)

func NewPostgresqlRepositories(db *sql.DB) *repos.Repositories {
	return &repos.Repositories{
		Example: NewExampleRepo(db),
	}
}
