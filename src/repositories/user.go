package repositories

import (
	"food-backend/src/domains"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetByEmail(email string) (user *domains.User, err error)
	SetRefreshToken(id string, refreshToken string) (err error)
	GetByRefreshToken(token string) (user *domains.User, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func (r userRepository) GetByEmail(email string) (user *domains.User, err error) {
	user = new(domains.User)

	sql := "SELECT * FROM users WHERE email = $1"
	err = r.db.Get(user, sql, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r userRepository) SetRefreshToken(id string, refreshToken string) (err error) {
	sql := "UPDATE users SET refresh_token = $1 WHERE user_id = $2"
	_, err = r.db.Exec(sql, refreshToken, id)
	if err != nil {
		return err
	}

	return nil
}

func (r userRepository) GetByRefreshToken(token string) (user *domains.User, err error) {
	user = new(domains.User)

	sql := "SELECT * from users WHERE refresh_token = $1"
	err = r.db.Get(user, sql, token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
