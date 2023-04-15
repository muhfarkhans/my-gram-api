package db

import (
	"log"
	"my-gram/comment"
	"my-gram/photo"
	"my-gram/socialmedia"
	"my-gram/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dbConfig postgres.Config) *gorm.DB {
	dbInstance, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	dbInstance.AutoMigrate(user.User{}, photo.Photo{}, comment.Comment{}, socialmedia.Socialmedia{})

	return dbInstance
}
