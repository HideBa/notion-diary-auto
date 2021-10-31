package database

import (
	"database/sql"
	"log"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(dbUrl string) error {
	m, err := migrate.New("file://infrastructure/database/migrations", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func NewDB(dbUrl string) *sql.DB {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("fail to connect with database")
	}
	return db
}

type RepositoryContainer struct {
	User domain.UserRepository
}

func NewRepository(db *sql.DB) *RepositoryContainer {
	return &RepositoryContainer{
		User: NewUserRepo(db),
	}
}
