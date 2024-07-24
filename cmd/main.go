package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dmarins/rate-limiter-challenge-go/config"
	"github.com/dmarins/rate-limiter-challenge-go/limiter"
	"github.com/dmarins/rate-limiter-challenge-go/middleware"
	"github.com/go-redis/redis/v8"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	rl := limiter.NewRateLimiter(
		rdb,
		cfg.RateLimitIP,
		cfg.RateLimitToken,
		time.Duration(cfg.BlockTimeIP)*time.Second,
		time.Duration(cfg.BlockTimeToken)*time.Second,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	handler := middleware.RateLimiterMiddleware(rl)(mux)

	http.ListenAndServe(":8080", handler)
}
