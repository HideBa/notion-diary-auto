package database

import (
	"database/sql"

	"github.com/HideBa/notion-diary-auto/domain"
	"github.com/HideBa/notion-diary-auto/pkg/id"
	"github.com/google/uuid"
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
	id       id.ID
	username string
}

func (ur *UserRepo) GetByID(id1 id.ID) (*domain.User, error) {
	var (
		uid      uuid.UUID
		username string
	)
	if err := ur.db.QueryRow("select * from users where id = ? LIMIT 1", id1.GetID()).Scan(&uid, &username); err != nil {
		return nil, err
	}
	um := &UserRepoModel{
		id:       id.FromIDUUID(uid),
		username: username,
	}
	return toUserModel(um), nil
}

func (ur *UserRepo) GetAll() ([]domain.User, error) {
	rows, err := ur.db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	users := []domain.User{}
	for rows.Next() {
		var (
			uid      uuid.UUID
			username string
		)
		err2 := rows.Scan(&uid, &username)
		if err2 != nil {
			return nil, err2
		}
		userRepoModel := UserRepoModel{
			id:       id.FromIDUUID(uid),
			username: username,
		}
		userDomainModel := toUserModel(&userRepoModel)
		users = append(users, *userDomainModel)
	}
	return users, nil
}

func (ur *UserRepo) Create(u *domain.User) (*domain.User, error) {
	return nil, nil
}
func (ur *UserRepo) Update(u *domain.User) error {
	return nil
}
func (ur *UserRepo) Delete(id id.ID) error {
	return nil
}

func toUserModel(u *UserRepoModel) *domain.User {
	return domain.NewUserWithId(u.id, u.username)
}
