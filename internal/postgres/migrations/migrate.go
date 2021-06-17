package migrations

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dataSourceName := "postgresql://postgres:postgres@localhost:5432/postgres"
	m, err := migrate.New(
		"file://internal/postgres/migrations",
		dataSourceName,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}