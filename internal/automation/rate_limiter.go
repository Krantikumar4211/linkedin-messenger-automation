package automation

import (
	"log"
	"time"
)

type RateLimiter struct {
	interval time.Duration
	lastTime time.Time
}

func NewRateLimiter(seconds int) *RateLimiter {
	return &RateLimiter{
		interval: time.Duration(seconds) * time.Second,
		lastTime: time.Now().Add(-time.Duration(seconds) * time.Second),
	}
}

func (r *RateLimiter) Wait() {
	elapsed := time.Since(r.lastTime)

	if elapsed < r.interval {
		waitTime := r.interval - elapsed
		log.Println("Rate limiter active, waiting:", waitTime)
		time.Sleep(waitTime)
	}

	r.lastTime = time.Now()
}
