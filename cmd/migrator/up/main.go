package main

import (
	"app/internal/config"
	"app/pkg/infra/database/postgresql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.MustLoad()

	m, err := migrate.New(
		"file://db/migrations",
		postgresql.CreateConnectionString(cfg),
	)
	if err != nil {
		fmt.Println("[Migrator] Failed read migrations")
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("[Migrator] No migrations to apply")
			return
		}

		fmt.Println("[Migrator] Failed up migrations")
		panic(err)
	}

	m.Close()
	fmt.Println("[Migrator] Migrations applied")
}
