package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

var (
	requests = make(map[string][]time.Time)
	mu       sync.Mutex
)

func RateLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		mu.Lock()
		now := time.Now()
		// Clean old requests
		var recent []time.Time
		for _, t := range requests[ip] {
			if now.Sub(t) < time.Minute {
				recent = append(recent, t)
			}
		}
		if len(recent) >= 100 {
			mu.Unlock()
			return c.Status(429).SendString("Too many requests")
		}
		requests[ip] = append(recent, now)
		mu.Unlock()
		return c.Next()
	}
}
