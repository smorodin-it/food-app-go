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

var DBCon *sqlx.DB

func ConnectToDB(config configType) error {
	db, err := sqlx.Connect(config.Driver,
		"user="+config.User+
			" password="+config.Password+
			" dbname="+config.DbName+
			" sslmode=disable")

	if err != nil {
		return err
	}

	DBCon = db

	return nil
}
