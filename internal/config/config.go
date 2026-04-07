package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	Port      string
	DBUrl     string
	JWTSecret string
	JWTExpire string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found! using system env.")
	}

	cfg := &Config{
		AppEnv:    getEnv("APP_ENV", "development"),
		Port:      getEnv("PORT", "8080"),
		DBUrl:     getEnv("DB_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "supersecretkey"),
		JWTExpire: getEnv("JWT_EXPIRE", "24h"),
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}
