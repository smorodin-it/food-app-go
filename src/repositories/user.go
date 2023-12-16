package repositories

import (
	"fmt"
	"food-backend/src/constants"
	"food-backend/src/database"
	"food-backend/src/domains"
	"github.com/google/uuid"
)

type UserRepository struct{}

func (u UserRepository) GetByEmail(email string) (*domains.User, error) {
	data := domains.User{}
	sql := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", constants.TableUsers)

	err := database.DBCon.Get(&data, sql, email)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (u UserRepository) SetRefreshToken(id string) (*string, error) {
	refreshToken := uuid.New().String()

	sql := fmt.Sprintf("UPDATE %s SET refresh_token = $1 WHERE id = $2", constants.TableUsers)

	_, err := database.DBCon.Exec(sql, refreshToken, id)
	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}
