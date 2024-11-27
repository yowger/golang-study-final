package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	MongoURI string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file: ", err)

	}

	return Config{
		AppPort:  os.Getenv("APP_PORT"),
		MongoURI: os.Getenv("MONGO_URI"),
	}
}
