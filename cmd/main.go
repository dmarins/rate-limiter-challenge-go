package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dmarins/rate-limiter-challenge-go/config"
	"github.com/dmarins/rate-limiter-challenge-go/middleware"
	"github.com/dmarins/rate-limiter-challenge-go/rl"
	"github.com/go-redis/redis/v8"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	var ratelimiter rl.RateLimiterInterface

	if cfg.Strategy == "redis" {
		rdb := redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddr,
		})

		ratelimiter = rl.NewRedisRateLimiter(
			rdb,
			cfg.RateLimitIP,
			cfg.RateLimitToken,
			time.Duration(cfg.BlockTimeIP)*time.Second,
			time.Duration(cfg.BlockTimeToken)*time.Second,
		)
	} else {
		ratelimiter = rl.NewInMemoryRateLimiter(
			cfg.RateLimitIP,
			cfg.RateLimitToken,
			time.Duration(cfg.BlockTimeIP)*time.Second,
			time.Duration(cfg.BlockTimeToken)*time.Second,
		)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	handler := middleware.RateLimiterMiddleware(ratelimiter)(mux)

	http.ListenAndServe(":8080", handler)
}
