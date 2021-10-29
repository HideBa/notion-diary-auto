package database

import (
	"database/sql"
	"log"

	"github.com/HideBa/notion-diary-auto/domain"
)

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
