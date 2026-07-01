package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Dsn  string
}

func LoadEnv() *Config {
	godotenv.Load()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .Env")
	}

	return &Config{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("DSN"),
	}

}
