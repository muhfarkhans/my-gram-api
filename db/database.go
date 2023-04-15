package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dbConfig postgres.Config) *gorm.DB {
	dbInstance, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	return dbInstance
}
