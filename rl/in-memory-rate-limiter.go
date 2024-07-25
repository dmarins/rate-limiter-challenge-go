package rl

import (
	"sync"
	"time"
)

type InMemoryRateLimiter struct {
	requests       map[string]int
	blocks         map[string]time.Time
	mutex          sync.Mutex
	rateLimitIP    int
	rateLimitToken int
	blockTimeIP    time.Duration
	blockTimeToken time.Duration
}

func NewInMemoryRateLimiter(rateLimitIP, rateLimitToken int, blockTimeIP, blockTimeToken time.Duration) RateLimiterInterface {
	return &InMemoryRateLimiter{
		requests:       make(map[string]int),
		blocks:         make(map[string]time.Time),
		rateLimitIP:    rateLimitIP,
		rateLimitToken: rateLimitToken,
		blockTimeIP:    blockTimeIP,
		blockTimeToken: blockTimeToken,
	}
}

func (imrl *InMemoryRateLimiter) Allow(ip, token string) (bool, error) {
	imrl.mutex.Lock()
	defer imrl.mutex.Unlock()

	if token != "" {
		return imrl.allowByToken(token)
	}

	return imrl.allowByIP(ip)
}

func (imrl *InMemoryRateLimiter) allowByIP(ip string) (bool, error) {
	key := "ip:" + ip

	return imrl.allow(key, imrl.rateLimitIP, imrl.blockTimeIP)
}

func (imrl *InMemoryRateLimiter) allowByToken(token string) (bool, error) {
	key := "token:" + token

	return imrl.allow(key, imrl.rateLimitToken, imrl.blockTimeToken)
}

func (imrl *InMemoryRateLimiter) allow(key string, limit int, blockTime time.Duration) (bool, error) {
	if blockedUntil, ok := imrl.blocks[key]; ok && time.Now().Before(blockedUntil) {
		return false, nil
	}

	if count, ok := imrl.requests[key]; ok {
		if count >= limit {
			imrl.blocks[key] = time.Now().Add(blockTime)
			return false, nil
		}

		imrl.requests[key] = count + 1
	} else {

		imrl.requests[key] = 1

		go imrl.resetCount(key)
	}

	return true, nil
}

func (imrl *InMemoryRateLimiter) resetCount(key string) {
	time.Sleep(time.Second)

	imrl.mutex.Lock()

	defer imrl.mutex.Unlock()

	delete(imrl.requests, key)
}
