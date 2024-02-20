package database

import "github.com/jmoiron/sqlx"

type configType struct {
	Driver   string
	User     string
	Password string
	DbName   string
}

var DbConfig = configType{
	Driver:   "postgres",
	User:     "food-app",
	Password: "food-app",
	DbName:   "food-app",
}

func Connect(config configType) (db *sqlx.DB, err error) {
	db, err = sqlx.Connect(config.Driver,
		"user="+config.User+
			" password="+config.Password+
			" dbname="+config.DbName+
			" sslmode=disable")

	if err != nil {
		return nil, err
	}

	return db, nil

}
