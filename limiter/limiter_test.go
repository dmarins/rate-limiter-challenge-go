package limiter

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestAllowByIp_WhenItDoesNotExceedsTheLimitOfRequestsPerIP(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root",
	})

	rl := NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	ip := "192.168.1.1"

	for i := 0; i < 5; i++ {
		allowed, err := rl.allowByIP(ctx, ip)

		assert.Nil(t, err)
		assert.True(t, allowed)
	}
}

func TestAllowByIp_WhenItExceedsTheLimitOfRequestsPerIP(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	ip := "192.168.1.1"

	for i := 0; i < 5; i++ {
		_, _ = rl.allowByIP(ctx, ip)
	}

	allowed, err := rl.allowByIP(ctx, ip)

	assert.Nil(t, err)
	assert.False(t, allowed)
}

func TestAllowByToken_WhenItDoesNotExceedsTheLimitOfRequestsPerToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	token := "abc123"

	for i := 0; i < 10; i++ {
		allowed, err := rl.allowByToken(ctx, token)

		assert.Nil(t, err)
		assert.True(t, allowed)
	}
}

func TestAllowByToken_WhenItExceedsTheLimitOfRequestsPerToken(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rl := NewRateLimiter(rdb, 5, 10, 5*time.Second, 5*time.Second)

	ctx := context.Background()
	rdb.FlushAll(ctx)

	token := "abc123"

	for i := 0; i < 10; i++ {
		_, _ = rl.allowByToken(ctx, token)
	}

	allowed, err := rl.allowByToken(ctx, token)

	assert.Nil(t, err)
	assert.False(t, allowed)
}
