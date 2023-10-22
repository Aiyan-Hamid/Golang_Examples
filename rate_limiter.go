package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens        int
	fillInterval  time.Duration
	capacity      int
	lastTimestamp time.Time
}

func NewRateLimiter(tokens int, fillInterval time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:        tokens,
		fillInterval:  fillInterval,
		capacity:      tokens,
		lastTimestamp: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	now := time.Now()
	elapsed := now.Sub(rl.lastTimestamp)
	tokensToAdd := int(elapsed.Seconds() / rl.fillInterval.Seconds())
	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		if rl.tokens > rl.capacity {
			rl.tokens = rl.capacity
		}
		rl.lastTimestamp = now
	}
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

func main() {
	limiter := NewRateLimiter(5, time.Second) // 5 tokens per second

	flag := false
	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Allowed request %d\n", i+1)
		} else {
			fmt.Printf("Rate limited request %d\n", i+1)
		}
		time.Sleep(200 * time.Millisecond) // Simulate requests every 200ms
		if i == 9 && flag == false {
			i = 5
			flag = true
			time.Sleep(time.Second)
		}
	}
}
