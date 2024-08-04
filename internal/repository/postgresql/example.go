package postgresql

import (
	"database/sql"
)

type ExampleRepo struct {
	db *sql.DB
}

func NewExampleRepo(db *sql.DB) *ExampleRepo {
	return &ExampleRepo{
		db: db,
	}
}

func (receiver ExampleRepo) Create(name string) {
	panic("IMPLEMENT ME...")
}

func (receiver ExampleRepo) Update() {
	panic("IMPLEMENT ME...")
}

func (receiver ExampleRepo) GetList() {
	panic("IMPLEMENT ME...")
}

func (receiver ExampleRepo) Delete() {
	panic("IMPLEMENT ME...")
}
