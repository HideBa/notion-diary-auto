package database

import (
	"database/sql"

	"github.com/HideBa/notion-diary-auto/domain"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) domain.UserRepository {
	return &UserRepo{
		db: db,
	}
}

type UserRepoModel struct {
	id       string
	username string
}

func (ur *UserRepo) GetByID(id string) (*domain.User, error) {
	var (
		id2      string
		username string
	)
	if err := ur.db.QueryRow("select * from users where id = ? LIMIT 1", id).Scan(&id2, &username); err != nil {
		return nil, err
	}
	um := &UserRepoModel{
		id:       id2,
		username: username,
	}
	return toUserModel(um), nil
}

func (ur *UserRepo) Create(u *domain.User) (*domain.User, error) {
	return nil, nil
}
func (ur *UserRepo) Update(u *domain.User) error {
	return nil
}
func (ur *UserRepo) Delete(id string) error {
	return nil
}

func toUserModel(u *UserRepoModel) *domain.User {
	return domain.NewUserWithId(u.id, u.username)
}
