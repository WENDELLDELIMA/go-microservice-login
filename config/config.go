package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env, usando variáveis do ambiente.")
	}

	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
}
