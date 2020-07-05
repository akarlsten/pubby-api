package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUser,
			Password: dbPass,
			Name:     dbName,
			Charset:  "utf8",
		},
	}
}
