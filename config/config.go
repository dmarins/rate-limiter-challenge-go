package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr      string
	RateLimitIP    int
	RateLimitToken int
	BlockTimeIP    int
	BlockTimeToken int
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	config := &Config{
		RedisAddr:      os.Getenv("REDIS_ADDR"),
		RateLimitIP:    getEnvAsInt("RATE_LIMIT_IP", 5),
		RateLimitToken: getEnvAsInt("RATE_LIMIT_TOKEN", 10),
		BlockTimeIP:    getEnvAsInt("BLOCK_TIME_IP", 300),
		BlockTimeToken: getEnvAsInt("BLOCK_TIME_TOKEN", 300),
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
