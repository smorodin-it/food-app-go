package repositories

import (
	"fmt"
	"food-backend/src/constants"
	"food-backend/src/database"
	"food-backend/src/domains"
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
