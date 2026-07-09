package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	GeminiAPIKey string
	QdrantURL    string
	QdrantAPIKey string
	UploadPath   string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Port:         getEnv("PORT", "8080"),
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),
		QdrantURL:    getEnv("QDRANT_URL", ""),
		QdrantAPIKey: getEnv("QDRANT_API_KEY", ""),
		UploadPath:   getEnv("UPLOAD_PATH", "uploads"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
