package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimiter struct {
	client         *redis.Client
	rateLimitIP    int
	rateLimitToken int
	blockTimeIP    time.Duration
	blockTimeToken time.Duration
}

func NewRateLimiter(client *redis.Client, rateLimitIP, rateLimitToken int, blockTimeIP, blockTimeToken time.Duration) *RateLimiter {
	return &RateLimiter{
		client:         client,
		rateLimitIP:    rateLimitIP,
		rateLimitToken: rateLimitToken,
		blockTimeIP:    blockTimeIP,
		blockTimeToken: blockTimeToken,
	}
}

func (rl *RateLimiter) Allow(ip, token string) (bool, error) {
	ctx := context.Background()
	if token != "" {
		return rl.AllowByToken(ctx, token)
	}
	return rl.AllowByIP(ctx, ip)
}

func (rl *RateLimiter) AllowByIP(ctx context.Context, ip string) (bool, error) {
	key := "rl:ip:" + ip
	return rl.allow(ctx, key, rl.rateLimitIP, rl.blockTimeIP)
}

func (rl *RateLimiter) AllowByToken(ctx context.Context, token string) (bool, error) {
	key := "rl:token:" + token
	return rl.allow(ctx, key, rl.rateLimitToken, rl.blockTimeToken)
}

func (rl *RateLimiter) allow(ctx context.Context, key string, limit int, blockTime time.Duration) (bool, error) {
	count, err := rl.client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		rl.client.Expire(ctx, key, time.Second)
	}

	if count > int64(limit) {
		rl.client.Set(ctx, key+":block", "1", blockTime)
		return false, nil
	}

	return true, nil
}
