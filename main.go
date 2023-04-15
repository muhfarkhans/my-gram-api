package main

import (
	"fmt"
	"my-gram/handler"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("app.config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
		panic(err)
	}

	router := gin.Default()
	handler.NewHandler(router)
	router.Run()
}
