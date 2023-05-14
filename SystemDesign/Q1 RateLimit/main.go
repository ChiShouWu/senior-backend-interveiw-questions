package main

import (
	"log"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiter struct {
	capacity    int
	rate        float64
	tokens      float64
	lastConsume time.Time
	mutex       sync.Mutex
}

func NewRateLimiter(capacity int, rate float64) *RateLimiter {
	return &RateLimiter{
		capacity:    capacity,
		rate:        rate,
		tokens:      float64(capacity),
		lastConsume: time.Now(),
	}
}

func (rl *RateLimiter) Consume() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastConsume)
	rl.lastConsume = now

	rl.tokens += float64(elapsed) * rl.rate
	if rl.tokens > float64(rl.capacity) {
		rl.tokens = float64(rl.capacity)
	}

	if rl.tokens < 1 {
		return false
	}

	rl.tokens -= 1
	return true
}
func main() {
	counter := 0

	// every 100ms to fill token
	limit := rate.Every(100 * time.Microsecond)
	// token bucket capacity = 5, every 100ms to fill
	rl := rate.NewLimiter(limit, 5)
	// set bucket capacity to 10
	rl.SetBurst(10)

	for {
		counter++

		// if token amount over 1 then take one token; otherwise reject
		// AllowN is thread safe method
		if isAllowed := rl.AllowN(time.Now(), 1); isAllowed {
			log.Printf("counter: %v, %v \n", counter, time.Now().Format(time.TimeOnly))
		} else {
			log.Printf("counter: %v, not allow \n", counter)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
