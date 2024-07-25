package rl

type RateLimiterInterface interface {
	Allow(ip, token string) (bool, error)
}
