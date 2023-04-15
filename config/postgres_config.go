package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
)

func PostgreConnectionLocal() string {
	dbConfigs := []string{
		"host=" + viper.GetString("database.host"),
		"dbname=" + viper.GetString("database.name"),
		"user=" + viper.GetString("database.username"),
		"password=" + viper.GetString("database.password"),
		"port=" + viper.GetString("database.port"),
		"sslmode=disable"}

	if os.Getenv("APP_ENV") == "production" {
		dbConfigs = []string{
			"host=" + os.Getenv("DB_HOST"),
			"dbname=" + os.Getenv("DB_NAME"),
			"user=" + os.Getenv("DB_USER"),
			"password=" + os.Getenv("DB_PASSWORD"),
			"port=" + os.Getenv("DB_PORT"),
			"sslmode=disable"}
	}

	return strings.Join(dbConfigs, " ")
}

func PostgreConfig() postgres.Config {
	configPostgre := postgres.Config{
		PreferSimpleProtocol: false,
	}

	configPostgre.DSN = PostgreConnectionLocal()

	return configPostgre
}
