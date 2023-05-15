package config

import (
	"os"
)

type DatabaseSerttings struct {
	DB       string
	User     string
	Password string
	Host     string
	Port     string
}

var DBnv DatabaseSerttings

func ParseDatabaseEnviroment() {
	DBnv = DatabaseSerttings{
		DB:       os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}
