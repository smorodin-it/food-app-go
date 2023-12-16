package repositories

import (
	"food-backend/src/database"
	"food-backend/src/domains"
)

type UserRepository struct{}

func (u UserRepository) GetByEmail(email string) (*domains.User, error) {
	data := new(domains.User)
	sql := "SELECT * FROM users WHERE email = $1"

	err := database.DBCon.Get(data, sql, email)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u UserRepository) SetRefreshToken(id string, refreshToken string) error {
	sql := "UPDATE users SET refresh_token = $1 WHERE id = $2"

	_, err := database.DBCon.Exec(sql, refreshToken, id)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetByRefreshToken(token string) (*domains.User, error) {
	data := new(domains.User)
	sql := "SELECT * from users WHERE refresh_token = $1"

	err := database.DBCon.Get(data, sql, token)
	if err != nil {
		return nil, err
	}

	return data, nil
}
