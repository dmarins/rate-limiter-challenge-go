package rl

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRateLimiter struct {
	client         *redis.Client
	rateLimitIP    int
	rateLimitToken int
	blockTimeIP    time.Duration
	blockTimeToken time.Duration
}

func NewRedisRateLimiter(client *redis.Client, rateLimitIP, rateLimitToken int, blockTimeIP, blockTimeToken time.Duration) RateLimiterInterface {
	return &RedisRateLimiter{
		client:         client,
		rateLimitIP:    rateLimitIP,
		rateLimitToken: rateLimitToken,
		blockTimeIP:    blockTimeIP,
		blockTimeToken: blockTimeToken,
	}
}

func (rl *RedisRateLimiter) Allow(ip, token string) (bool, error) {
	ctx := context.Background()

	if token != "" {
		return rl.allowByToken(ctx, token)
	}

	return rl.allowByIP(ctx, ip)
}

func (rl *RedisRateLimiter) allowByIP(ctx context.Context, ip string) (bool, error) {
	key := "rrl:ip:" + ip

	return rl.allow(ctx, key, rl.rateLimitIP, rl.blockTimeIP)
}

func (rl *RedisRateLimiter) allowByToken(ctx context.Context, token string) (bool, error) {
	key := "rrl:token:" + token

	return rl.allow(ctx, key, rl.rateLimitToken, rl.blockTimeToken)
}

func (rl *RedisRateLimiter) allow(ctx context.Context, key string, limit int, blockTime time.Duration) (bool, error) {
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
