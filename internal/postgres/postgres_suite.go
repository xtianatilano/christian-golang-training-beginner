package postgres

import (
	"database/sql"

	// This is imported for migrations
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/stretchr/testify/suite"
)

const (
	migrationDbName = "postgres"

	// DefaultTestDsn is the default url for testing postgresql in the postgres test suites
	DefaultTestDsn = "user=postgres password=password dbname=testing host=localhost port=5432 sslmode=disable"
)

// Suite struct for MySQL Suite
type Suite struct {
	suite.Suite
	DSN                     string
	DBConn                  *sql.DB
	Migration               *migration
	MigrationLocationFolder string
	DBName                  string
}

// SetupSuite setup at the beginning of test
func (s *Suite) SetupSuite() {
	var err error
	s.DBConn, err = sql.Open(migrationDbName, s.DSN)
	s.Require().NoError(err)
	err = s.DBConn.Ping()
	s.Require().NoError(err)
	s.Migration, err = runMigration(s.DBConn, s.MigrationLocationFolder)
	s.Require().NoError(err)
}

// TearDownSuite teardown at the end of test
func (s *Suite) TearDownSuite() {
	err := s.DBConn.Close()
	s.Require().NoError(err)
}