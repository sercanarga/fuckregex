package db

import (
	"fmt"
	"fuckregex/model/db_model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dbConfig *db_model.Config) error {
	var err error

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port, dbConfig.SSLMode),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return err
	}
	return nil
}
