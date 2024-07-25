package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr      string
	RateLimitIP    int
	RateLimitToken int
	BlockTimeIP    int
	BlockTimeToken int
	Strategy       string
}

func LoadConfig() (*Config, error) {
	// Get the directory of the executable or the working directory
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(ex)

	// Determine the absolute path to the .env file
	envPath := filepath.Join(exeDir, "..", ".env")

	// Load the .env file
	err = godotenv.Load(envPath)
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{
		RedisAddr:      os.Getenv("REDIS_ADDR"),
		RateLimitIP:    getEnvAsInt("RATE_LIMIT_IP", 5),
		RateLimitToken: getEnvAsInt("RATE_LIMIT_TOKEN", 10),
		BlockTimeIP:    getEnvAsInt("BLOCK_TIME_IP", 300),
		BlockTimeToken: getEnvAsInt("BLOCK_TIME_TOKEN", 300),
		Strategy:       os.Getenv("STRATEGY"),
	}

	return config, nil
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if valueStr == "" {
		return defaultVal
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultVal
	}

	return value
}
