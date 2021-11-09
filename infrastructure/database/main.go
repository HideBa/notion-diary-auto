package database

import (
	"database/sql"
	"log"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(driver database.Driver, migrateVer uint) error {
	m, err := migrate.NewWithDatabaseInstance("file://infrastructure/database/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	migrateV, dirty, err := m.Version()
	if dirty {
		log.Fatal("DB is dirty")
	}
	if migrateV != migrateVer || err != nil {
		if err := m.Migrate(migrateVer); err != nil {
			log.Fatal(err)
		}
		return nil
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
