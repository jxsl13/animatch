package common

import (
	"time"

	"golang.org/x/time/rate"
)

func NewRateLimiter(every time.Duration, calls int) *rate.Limiter {
	return rate.NewLimiter(rate.Every(every/time.Duration(calls)), 1)
}
