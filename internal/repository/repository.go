package repository

import (
	"app/internal/repository/postgresql"
	"database/sql"
)

type IExampleRepo interface {
	Create(name string)
	Update()
	Delete()
	GetList()
}

type Repositories struct {
	Example IExampleRepo
}

func NewRepositories(db *sql.DB) *Repositories {
	return postgresql.NewPostgresqlRepositories(db)
}
