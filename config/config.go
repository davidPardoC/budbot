package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Telegram Telegram
	Database Database
	Server   Server
}

type Telegram struct {
	Token string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Server struct {
	Port string
	Env  string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Telegram: Telegram{
			Token: os.Getenv("TELEGRAM_TOKEN"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Server: Server{
			Port: os.Getenv("PORT"),
			Env:  os.Getenv("ENV"),
		},
	}
}
