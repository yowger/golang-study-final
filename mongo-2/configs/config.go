package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	MongoURI string
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, errors.New("failed to load .env file")
	}

	return Config{
		AppPort:  os.Getenv("APP_PORT"),
		MongoURI: os.Getenv("MONGO_URI"),
	}, nil
}
