package repository

import (
	"app/internal/repository/postgresql"
	"app/internal/repository/repos"
	"database/sql"
)

func NewRepositories(db *sql.DB) *repos.Repositories {
	return postgresql.NewPostgresqlRepositories(db)
}
