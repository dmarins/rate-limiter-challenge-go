package limiter_test

import (
	"context"
	"testing"
	"time"

	"github.com/dmarins/rate-limiter-challenge-go/limiter"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestAllowByIp_WhenItDoesNotExceedsTheLimitOfRequestsPerIP(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := limiter.NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	ip := "192.168.1.1"

	for i := 0; i < 5; i++ {
		allowed, err := rl.AllowByIP(ctx, ip)

		assert.Nil(t, err)
		assert.True(t, allowed)
	}
}

func TestAllowByIp_WhenItExceedsTheLimitOfRequestsPerIP(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := limiter.NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	ip := "192.168.1.1"

	for i := 0; i < 5; i++ {
		_, _ = rl.AllowByIP(ctx, ip)
	}

	allowed, err := rl.AllowByIP(ctx, ip)

	assert.Nil(t, err)
	assert.False(t, allowed)
}

func TestAllowByToken_WhenItDoesNotExceedsTheLimitOfRequestsPerToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := limiter.NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	token := "abc123"

	for i := 0; i < 10; i++ {
		allowed, err := rl.AllowByToken(ctx, token)

		assert.Nil(t, err)
		assert.True(t, allowed)
	}
}

func TestAllowByToken_WhenItExceedsTheLimitOfRequestsPerToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := limiter.NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	token := "abc123"

	for i := 0; i < 10; i++ {
		_, _ = rl.AllowByToken(ctx, token)
	}

	allowed, err := rl.AllowByToken(ctx, token)

	assert.Nil(t, err)
	assert.False(t, allowed)
}
